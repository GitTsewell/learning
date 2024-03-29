# mutex
```
type Mutex struct {
	state int32
	sema  uint32
}
```

![sema](./sync_sema.png)

在默认情况下，互斥锁的所有状态位都是 0，int32 中的不同位分别表示了不同的状态：
互斥锁的状态
+ mutexLocked — 表示互斥锁的锁定状态
+ mutexWoken — 表示从正常模式被从唤醒
+ mutexStarving — 当前的互斥锁进入饥饿状态
+ waitersCount — 当前互斥锁上等待的 Goroutine 个数

## mutex lock()
1. 首先如果当前锁处于初始化状态就直接用 CAS (atomic.CompareAndSwapInt32()) 方法尝试获取锁
```
if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		if race.Enabled {
			race.Acquire(unsafe.Pointer(m))
		}
		return
	}
```
2. 如果获取锁失败就进入 lockSlow()方法
3. 会首先判断当前能不能进入自旋状态，如果可以就进入自旋，最多自旋 4 次
```
if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			// Active spinning makes sense.
			// Try to set mutexWoken flag to inform Unlock
			// to not wake other blocked goroutines.
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
				atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				awoke = true
			}
			runtime_doSpin()
			iter++
			old = m.state
			continue
		}
```
4. 自旋完成之后，就会去计算当前的锁的状态
5. 然后尝试通过 CAS 获取锁
6. 如果没有获取到就调用 runtime_SemacquireMutex 方法休眠当前 goroutine 并且尝试获取信号量
```
runtime_SemacquireMutex(&m.sema, queueLifo, 1)
```
7. goroutine 被唤醒之后会先判断当前是否处在饥饿状态，（如果当前 goroutine 超过 1ms 都没有获取到锁就会进饥饿模式）
```
if old&mutexStarving != 0 {}
```
8. 如果处在饥饿状态就会获得互斥锁，如果等待队列中只存在当前 Goroutine，互斥锁还会从饥饿模式中退出 
```
if old&mutexStarving != 0 {
    if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
        throw("sync: inconsistent mutex state")
    }
    delta := int32(mutexLocked - 1<<mutexWaiterShift)
    if !starving || old>>mutexWaiterShift == 1 {
        delta -= mutexStarving
    }
    atomic.AddInt32(&m.state, delta)
    break
}
```
9. 如果不在，就会设置唤醒和饥饿标记、重置迭代次数并重新执行获取锁的循环
```
awoke = true
iter = 0
```

## unlock()
1. 直接调用 atomic.AddInt32()进行解锁,成功直接结束.失败调用unlockSlow()函数.解锁一个没有锁定的互斥量会报运行时错误
2. 判断是否处于饥饿模式
3. 饥饿模式，走 handoff 流程，直接将锁交给下一个等待的 goroutine，注意这个时候不会从饥饿模式中退出
```
runtime_Semrelease(&m.sema, true, 1)
```
4. 正常模式下,如果当前没有等待者.或者 goroutine 已经被唤醒或者是处于锁定状态了，就直接返回
```
if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
    return
}
```
5. 如果有等待者唤醒等待者并且移交锁的控制权
```
new = (old - 1<<mutexWaiterShift) | mutexWoken
if atomic.CompareAndSwapInt32(&m.state, old, new) {
    runtime_Semrelease(&m.sema, false, 1)
    return
}
old = m.state
```

# RWMutex
```
type RWMutex struct {
	w           Mutex  // 复用互斥锁
	writerSem   uint32 // 信号量，用于写等待读
	readerSem   uint32 // 信号量，用于读等待写
	readerCount int32  // 当前执行读的 goroutine 数量
	readerWait  int32  // 写操作被阻塞的准备读的 goroutine 的数量
}
```

## 读锁
### RLock()
```
if atomic.AddInt32(&rw.readerCount, 1) < 0 {
    // A writer is pending, wait for it.
    runtime_SemacquireMutex(&rw.readerSem, false, 0)
}
```
首先是读锁， atomic.AddInt32(&rw.readerCount, 1)  调用这个原子方法，对当前在读的数量加一，如果返回负数，那么说明当前有其他写锁，这时候就调用 runtime_SemacquireMutex  休眠 goroutine 等待被唤醒

### RUnlock()
```
func (rw *RWMutex) RUnlock() {
	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
		// Outlined slow-path to allow the fast-path to be inlined
		rw.rUnlockSlow(r)
	}
}
```
解锁的时候对正在读的操作减一，如果返回值小于 0 那么说明当前有在写的操作，这个时候调用 rUnlockSlow  进入慢速通道
```
func (rw *RWMutex) rUnlockSlow(r int32) {
	if r+1 == 0 || r+1 == -rwmutexMaxReaders {
		race.Enable()
		throw("sync: RUnlock of unlocked RWMutex")
	}
	// A writer is pending.
	if atomic.AddInt32(&rw.readerWait, -1) == 0 {
		// The last reader unblocks the writer.
		runtime_Semrelease(&rw.writerSem, false, 1)
	}
}
```
被阻塞的准备读的 goroutine 的数量减一，readerWait 为 0，就表示当前没有正在准备读的 goroutine 这时候调用 runtime_Semrelease  唤醒写操作

## 写锁
### Lock()
```
func (rw *RWMutex) Lock() {
	// First, resolve competition with other writers.
	rw.w.Lock()
	// Announce to readers there is a pending writer.
	r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
	// Wait for active readers.
	if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
		runtime_SemacquireMutex(&rw.writerSem, false, 0)
	}
}
```
1. 首先调用互斥锁的 lock，获取到互斥锁之后，
2. atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders)  调用这个函数阻塞后续的读操作
3. 如果计算之后当前仍然有其他 goroutine 持有读锁，那么就调用 runtime_SemacquireMutex  休眠当前的 goroutine 等待所有的读操作完成

### UnLock()
```
func (rw *RWMutex) Unlock() {
	// Announce to readers there is no active writer.
	r := atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)
	if r >= rwmutexMaxReaders {
		race.Enable()
		throw("sync: Unlock of unlocked RWMutex")
	}
	// Unblock blocked readers, if any.
	for i := 0; i < int(r); i++ {
		runtime_Semrelease(&rw.readerSem, false, 0)
	}
}
```
解锁的操作，会先调用 atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)  将恢复之前写入的负数，然后根据当前有多少个读操作在等待，循环唤醒

# WaitGroup
```
type WaitGroup struct {
	noCopy noCopy

	// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
	// 64-bit atomic operations require 64-bit alignment, but 32-bit
	// compilers do not ensure it. So we allocate 12 bytes and then use
	// the aligned 8 bytes in them as state, and the other 4 as storage
	// for the sema.
	state1 [3]uint32
}
```
WaitGroup 结构十分简单，由 nocopy 和 state1 两个字段组成，其中 nocopy 是用来防止复制的
```
type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}
```
sync.noCopy 是一个特殊的私有结构体，tools/go/analysis/passes/copylock 包中的分析器会在编译期间检查被拷贝的变量中是否包含 sync.noCopy 
或者实现了 Lock 和 Unlock 方法，如果包含该结构体或者实现了对应的方法就会报出以下错误：

state1 的设计非常巧妙，这是一个是十二字节的数据，这里面主要包含两大块，counter 占用了 8 字节用于计数，sema 占用 4 字节用做信号量

## ADD
Add 其实最主要的就是加上计数器的值
```
func (wg *WaitGroup) Add(delta int) {
    // 先从 state 当中把数据和信号量取出来
	statep, semap := wg.state()

    // 在 waiter 上加上 delta 值
	state := atomic.AddUint64(statep, uint64(delta)<<32)
    // 取出当前的 counter
	v := int32(state >> 32)
    // 取出当前的 waiter，正在等待 goroutine 数量
	w := uint32(state)

    // counter 不能为负数
	if v < 0 {
		panic("sync: negative WaitGroup counter")
	}

    // 这里属于防御性编程
    // w != 0 说明现在已经有 goroutine 在等待中，说明已经调用了 Wait() 方法
    // 这时候 delta > 0 && v == int32(delta) 说明在调用了 Wait() 方法之后又想加入新的等待者
    // 这种操作是不允许的
	if w != 0 && delta > 0 && v == int32(delta) {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
    // 如果当前没有人在等待就直接返回，并且 counter > 0
	if v > 0 || w == 0 {
		return
	}

    // 这里也是防御 主要避免并发调用 add 和 wait
	if *statep != state {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}

	// 唤醒所有 waiter，看到这里就回答了上面的问题了
	*statep = 0
	for ; w != 0; w-- {
		runtime_Semrelease(semap, false, 0)
	}
}
``` 

## Wait
wait 主要就是等待其他的 goroutine 完事之后唤醒
```
func (wg *WaitGroup) Wait() {
	// 先从 state 当中把数据和信号量的地址取出来
    statep, semap := wg.state()

	for {
     	// 这里去除 counter 和 waiter 的数据
		state := atomic.LoadUint64(statep)
		v := int32(state >> 32)
		w := uint32(state)

        // counter = 0 说明没有在等的，直接返回就行
        if v == 0 {
			// Counter is 0, no need to wait.
			return
		}

		// waiter + 1，调用一次就多一个等待者，然后休眠当前 goroutine 等待被唤醒
		if atomic.CompareAndSwapUint64(statep, state, state+1) {
			runtime_Semacquire(semap)
			if *statep != 0 {
				panic("sync: WaitGroup is reused before previous Wait has returned")
			}
			return
		}
	}
}
```
## Done
Done 只是 Add 的简单封装，所以实际上是可以通过一次加一个比较大的值减少调用，或者达到快速唤醒的目的。
```
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}
```

# Sync.Once
```
type Once struct {
	done uint32
	m    Mutex
}
```
done == 0 表示每执行过,m就是一个互斥锁

## do()
```
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}
```
如果done==0 代表没执行过,调用doSlow()
```
func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```
doSlow 很简单就是先上锁 ,然后执行完f方法之后把done+1

# errGroup
```
type Group struct {
    // context 的 cancel 方法
	cancel func()

    // 复用 WaitGroup
	wg sync.WaitGroup

	// 用来保证只会接受一次错误
	errOnce sync.Once
    // 保存第一个返回的错误
	err     error
}
```

## withContext
```
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}
```
WithContext  就是使用 WithCancel  创建一个可以取消的 context 将 cancel 赋值给 Group 保存起来，然后再将 context 返回回去

## Go
```
func (g *Group) Go(f func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}
```
Go  方法其实就类似于 go  关键字，会启动一个携程，然后利用 waitgroup  来控制是否结束，如果有一个非 nil  的 error 出现就会保存起来并且如果有 cancel
就会调用 cancel  取消掉，使 ctx  返回

## wait
```
func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
```
Wait  方法其实就是调用 WaitGroup  等待，如果有 cancel  就调用一下

# Semaphore
golang 官方的带权重的信号量控制机制
```
type Weighted struct {
    size    int64 // 设置一个最大权值
    cur     int64 // 标识当前已被使用的资源数
    mu      sync.Mutex // 提供临界区保护
    waiters list.List // 阻塞等待的调用者列表
}
```

这个结构体对外暴露四个方法
1. golang/sync/semaphore.NewWeighted 用于创建新的信号量；
2. golang/sync/semaphore.Weighted.Acquire 阻塞地获取指定权重的资源，如果当前没有空闲资源，会陷入休眠等待；
3. golang/sync/semaphore.Weighted.TryAcquire 非阻塞地获取指定权重的资源，如果当前没有空闲资源，会直接返回 false；
4. golang/sync/semaphore.Weighted.Release 用于释放指定权重的资源；

其实原理很简单,首先根据size,和cur判断当前权重,Acquire 和 TryAcquire 获取权重资源就设计锁和cur的增加 ,acquire没获取到指定权重阻塞是加了一个channel,然后select监听等待
在waiters列表里面 ,调用release释放的时候,在cur减去权重,在判断waiters里面是否有等待的,如果有调用close触发channel

# singleflight
这个库的主要作用就是将一组相同的请求合并成一个请求，实际上只会去请求一次，然后对所有的请求返回相同的结果。

这个结构体对外暴露三个方法
```
type Group
    // Do 执行函数, 对同一个 key 多次调用的时候，在第一次调用没有执行完的时候
	// 只会执行一次 fn 其他的调用会阻塞住等待这次调用返回
	// v, err 是传入的 fn 的返回值
	// shared 表示是否真正执行了 fn 返回的结果，还是返回的共享的结果
    func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool)

	// DoChan 和 Do 类似，只是 DoChan 返回一个 channel，也就是同步与异步的区别
	func (g *Group) DoChan(key string, fn func() (interface{}, error)) <-chan Result

    // Forget 用于通知 Group 删除某个 key 这样后面继续这个 key 的调用的时候就不会在阻塞等待了
	func (g *Group) Forget(key string)
```

结构
```
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

type call struct {
	wg sync.WaitGroup

	// 函数的返回值，在 wg 返回前只会写入一次
	val interface{}
	err error

	// 使用调用了 Forgot 方法
	forgotten bool

    // 统计调用次数以及返回的 channel
	dups  int
	chans []chan<- Result
}
```

## Do
```
func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
	g.mu.Lock()

    // 前面提到的懒加载
    if g.m == nil {
		g.m = make(map[string]*call)
	}

    // 会先去看 key 是否已经存在
	if c, ok := g.m[key]; ok {
       	// 如果存在就会解锁
		c.dups++
		g.mu.Unlock()

        // 然后等待 WaitGroup 执行完毕，只要一执行完，所有的 wait 都会被唤醒
		c.wg.Wait()

        // 这里区分 panic 错误和 runtime 的错误，避免出现死锁，后面可以看到为什么这么做
		if e, ok := c.err.(*panicError); ok {
			panic(e)
		} else if c.err == errGoexit {
			runtime.Goexit()
		}
		return c.val, c.err, true
	}

    // 如果我们没有找到这个 key 就 new call
	c := new(call)

    // 然后调用 waitgroup 这里只有第一次调用会 add 1，其他的都会调用 wait 阻塞掉
    // 所以这要这次调用返回，所有阻塞的调用都会被唤醒
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

    // 然后我们调用 doCall 去执行
	g.doCall(c, key, fn)
	return c.val, c.err, c.dups > 0
}
```
原理很简单,先加锁,然后判断key是否在map中有值,如果有就waitGroup等待执行,如果没有就add(1)阻塞执行,最后解锁,然后调用call()

# sync map
结构模型
```
type Map struct {
    mu Mutex
    read atomic.Value // readOnly
    dirty map[interface{}]*entry
    misses int
}
```
```
type readOnly struct {
	m       map[interface{}]*entry
	amended bool // true if the dirty map contains some key not in m.
}
```

两个map,优先存dirty中,触发一定条件后刷进read中

每次操作dirty都有mutex锁

![sync_map](./sync_map.png)

流程概述:
1. 读的时候先去readonly里面找,找不到就判断amended,如果为true,说明dirty里面可能有,就去dirty里面找,然后....
2. 写的时候先判断readonly是否有,有直接更新,没有就去dirty里面新建一个
3. 相信流程看下面文章链接

! [年度最佳【golang】sync.Map详解](https://segmentfault.com/a/1190000023879083#item-4-2)

# sync pool
! [年度最佳【golang】sync.Pool详解](https://segmentfault.com/a/1190000023878185)

