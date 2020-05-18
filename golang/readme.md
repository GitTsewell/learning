<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-generate-toc again -->

* [defer和命名返回值](#defer和命名返回值)
* [defer的执行顺序](#defer的执行顺序)
* [make和new的区别](#make和new的区别)
* [goroutine的特点](#goroutine的特点)
* [channel的特点](#channel的特点)
* [context包](#context包)
* [反射](#反射)
* [锁](#锁)
  * [互斥锁 sync.mutex](#互斥锁-syncmutex)
  * [读写锁 sync.rwmutex](#读写锁-syncrwmutex)
  * [sync.waitGroup](#syncwaitGroup)
  * [sync.once](#synconce)


<!-- markdown-toc end -->

## defer和命名返回值
当函数有可命名结果形参是,结果形参在调用的时候就被初始化设置为零值  
当函数有可命名结果形参时，defer函数是可以修改它，然后再将它的值返回  
函数没有可命名结果形参，t只是个普通局部变量，defer无法对返回值做修改
```
func main() {
	fmt.Println(deferFunc1(1))
	fmt.Println(deferFunc2(1))
	fmt.Println(deferFunc3(1))
}

func deferFunc1(i int) (t int) {
	t = i
	defer func() {
		x += 3
	}()

	return
}

func deferFunc2(i int) int {
	t = i
	defer func() {
		x += 3
	}()

	return t
}

func deferFunc1(i int) (t int) {
	defer func() {
		x += i
	}()

	return 2
}
```
输出 4,1,3

## defer的执行顺序
Go的 defer 语句用于预设一个函数调用（即推迟执行函数）， 该函数会在执行 defer 的函数返回之前立即执行
```
for i := 0; i < 5; i++ {
	defer fmt.Printf("%d ", i)
}
```
被推迟的函数按照后进先出（LIFO）的顺序执行，因此以上代码在函数返回时会打印 4 3 2 1 0

## make和new的区别
make和new都是为变量分配内存空间
+ make返回值,new返回指针
+ make只能作用于切片,map和channel

## goroutine的特点
+ 和其他多线程相比,避免内核态和用户态的切换导致的成本(基本切换过来就可以用,而不用做一些清理工作)
+ 由语言层面进行调度,只要遇到io阻塞就可以切换下一个等待执行的goroutine,不会让cpu空等(类似nodejs的epoll)
+ 更小的栈空间允许创建大量的实例(一个一般goroutine大小在2~4k)  
goroutine的调度模型是MPG  cpu,线程,g对象(goroutine),cpu调度,线程是真正干活的地方,g排队等待  
goroutine一定要有return或者中断,不然会造成内存泄露

## channel的特点
go使用channel来实现共享数据,而不是通过共享内存来实现  
channel分为有缓冲和无缓冲

## context包
context.context是一个golang1.7引入标准库的接口,该接口定义四个需要实现的方法
+ deadline 完成工作的截止时间
+ done 返回一个channel,多次调用返回同一个channel
+ err context.context结束的原因
+ value context.context中获取键对应的值
context.contrxt的作用,goroutine构成的树形结构中对信号进行同步以减少计算资源的浪费,在不同的goroutine之间同步请求特定数据,取消信号以及处理请求的截止日期

## 反射
运行时反射是程序在运行期间检查其自身结构的一种方式,获取到一个程序的值和类型,那么久意味着知道了这个变量的全部信息,反射在平时中并不常用,但是也有一些适合的应用场景
+ 动态调用函数 ```reflect.ValueOf(t).MethodByName(name1).Call(nil)```
+ struct tag 解析 这个应该是平时是最常用的 ```field.Tag.Lookup("json")  field.Tag.Get("test") ```
+ 类型转换与赋值 ```newTv.Elem().FieldByName(newTTag).Set(tValue)```
+ 通过kind()处理不同分支 ```t := reflect.TypeOf(a) switch t.Kind() ```
+ 判断实例是否实现了某接口  
[参考链接](https://blog.csdn.net/ajfgurjfmvvlsfkjglkh/article/details/85637417)

## 锁
### 互斥锁 sync.mutex
在默认情况下，互斥锁的所有状态位都是 0，int32 中的不同位分别表示了不同的状态：
+ mutexLocked — 表示互斥锁的锁定状态；
+ mutexWoken — 表示从正常模式被从唤醒；
+ mutexStarving — 当前的互斥锁进入饥饿状态；
+ waitersCount — 当前互斥锁上等待的 Goroutine 个数  

#### 正常模式和饥饿模式
sync.Mutex 有两种模式 — 正常模式和饥饿模式。我们需要在这里先了解正常模式和饥饿模式都是什么，它们有什么样的关系。  
在正常模式下，锁的等待者会按照先进先出的顺序获取锁。但是刚被唤起的 Goroutine 与新创建的 Goroutine 竞争时，大概率会获取不到锁，为了减少这种情况的出现，一旦 Goroutine 超过 1ms 没有获取到锁，它就会将当前互斥锁切换饥饿模式，防止部分 Goroutine 被『饿死』
饥饿模式是在 Go 语言 1.9 版本引入的优化1，引入的目的是保证互斥锁的公平性（Fairness）。

在饥饿模式中，互斥锁会直接交给等待队列最前面的 Goroutine。新的 Goroutine 在该状态下不能获取锁、也不会进入自旋状态，它们只会在队列的末尾等待。如果一个 Goroutine 获得了互斥锁并且它在队列的末尾或者它等待的时间少于 1ms，那么当前的互斥锁就会被切换回正常模式。
相比于饥饿模式，正常模式下的互斥锁能够提供更好地性能，饥饿模式的能避免 Goroutine 由于陷入等待无法获取锁而造成的高尾延时。

#### 小结
**互斥锁的加锁过程比较复杂，它涉及自旋、信号量以及调度等概念：**

+ 如果互斥锁处于初始化状态，就会直接通过置位 mutexLocked 加锁；
+ 如果互斥锁处于 mutexLocked 并且在普通模式下工作，就会进入自旋，执行 30 次 PAUSE 指令消耗 CPU 时间等待锁的释放；
+ 如果当前 Goroutine 等待锁的时间超过了 1ms，互斥锁就会切换到饥饿模式；
+ 互斥锁在正常情况下会通过 sync.runtime_SemacquireMutex 函数将尝试获取锁的 Goroutine 切换至休眠状态，等待锁的持有者唤醒当前 Goroutine；
+ 如果当前 Goroutine 是互斥锁上的最后一个等待的协程或者等待的时间小于 1ms，当前 Goroutine 会将互斥锁切换回正常模式；


**互斥锁的解锁过程与之相比就比较简单，其代码行数不多、逻辑清晰，也比较容易理解:**

+ 当互斥锁已经被解锁时，那么调用 sync.Mutex.Unlock 会直接抛出异常；
+ 当互斥锁处于饥饿模式时，会直接将锁的所有权交给队列中的下一个等待者，等待者会负责设置 mutexLocked 标志位；
+ 当互斥锁处于普通模式时，如果没有 Goroutine 等待锁的释放或者已经有被唤醒的 Goroutine 获得了锁，就会直接返回；在其他情况下会通过 sync.runtime_Semrelease 唤醒对应的 Goroutine；

### 读写锁 sync.rwmutex
写操作使用 sync.RWMutex.Lock 和 sync.RWMutex.Unlock 方法；
读操作使用 sync.RWMutex.RLock 和 sync.RWMutex.RUnlock 方法

-|读|写
--|:--:|--:
读|Y|N
写|N|N

#### 小结
+ 调用 sync.RWMutex.Lock 尝试获取写锁时；
    + 每次 sync.RWMutex.RUnlock 都会将 readerWait 其减一，当它归零时该 Goroutine 就会获得写锁；
    + 将 readerCount 减少 rwmutexMaxReaders 个数以阻塞后续的读操作；
+ 调用 sync.RWMutex.Unlock 释放写锁时，会先通知所有的读操作，然后才会释放持有的互斥锁

### sync.waitGroup
sync.waitGroup可以等待一组goroutine的返回,一个比较常见的使用场景是批量发出rpc或http请求

+ sync.WaitGroup 必须在 sync.WaitGroup.Wait 方法返回之后才能被重新使用；
+ sync.WaitGroup.Done 只是对 sync.WaitGroup.Add 方法的简单封装，我们可以向 sync.WaitGroup.Add 方法传入任意负数（需要保证计数器非负）快速将计数器归零以唤醒其他等待的 Goroutine；
+ 可以同时有多个 Goroutine 等待当前 sync.WaitGroup 计数器的归零，这些 Goroutine 会被同时唤醒

### sync.once
sync.once可以保证在go程序运行期间的某段代码只会执行一次