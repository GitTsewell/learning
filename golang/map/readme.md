# Map

## 数据结构

### hmap
```
type hmap struct {
    count     int 
    flags     uint8
    B         uint8
    noverflow uint16
    hash0     uint32
    buckets    unsafe.Pointer
    oldbuckets unsafe.Pointer
    nevacuate  uintptr 
    extra *mapextra
}

type mapextra struct {
    overflow    *[]*bmap
    oldoverflow *[]*bmap
    nextOverflow *bmap
}
```
+ count：map 的大小，也就是 len() 的值。代指 map 中的键值对个数
+ flags：状态标识，主要是 goroutine 写入和扩容机制的相关状态控制。并发读写的判断条件之一就是该值
+ B：桶，最大可容纳的元素数量，值为 负载因子（默认 6.5） * 2 ^ B，是 2 的指数
+ noverflow：溢出桶的数量
+ hash0：哈希因子
+ buckets：保存当前桶数据的指针地址（指向一段连续的内存地址，主要存储键值对数据
+ oldbuckets，保存旧桶的指针地址
+ nevacuate：迁移进度 如果正在扩容,标明接下来该扩容的桶
+ 原有 buckets 满载后，会发生扩容动作，在 Go 的机制中使用了增量扩容，如下为细项
  + overflow 为 hmap.buckets （当前）溢出桶的指针地址
  + oldoverflow 为 hmap.oldbuckets （旧）溢出桶的指针地址
  + nextOverflow 为空闲溢出桶的指针地址

### bmap
```
type bmap struct {
	tophash [bucketCnt]uint8
}
```

#### tophash
tophash 是个长度为 8 的数组，代指桶最大可容纳的键值对为 8。
存储每个元素 hash 值的高 8 位，如果 tophash [0] <minTopHash(5)，则 tophash [0] 表示为迁移进度

在运行期间，runtime.bmap 结构体其实不止包含 tophash 字段，因为哈希表中可能存储不同类型的键值对，而且 Go 语言也不支持泛型，所以键值对占据的
内存空间大小只能在编译时进行推导

```
  type bmap struct {
    topbits  [8]uint8
    keys     [8]keytype
    values   [8]valuetype
    pad      uintptr
    overflow uintptr
}
```

#### keys 和 values
在这里我们留意到，存储 k 和 v 的载体并不是用 k/v/k/v/k/v/k/v 的模式，而是 k/k/k/k/v/v/v/v 的形式去存储
```
map[int64]int8
```
在这个例子中，如果按照 k/v/k/v/k/v/k/v 的形式存放的话，虽然每个键值对的值都只占用 1 个字节。但是却需要 7 个填充字节来补齐内存空间。最终就会造成大量的内存 “浪费”

但是如果以 k/k/k/k/v/v/v/v 的形式存放的话，就能够解决因对齐所 "浪费" 的内存空间 因此这部分的拆分主要是考虑到内存对齐的问题，虽然相对会复杂一点，但依然值得如此设计

#### overflow
而在 Go map 中当 hmap.buckets 满了后，就会使用溢出桶接着存储。

## 访问
+ 当接受一个参数时，会使用 runtime.mapaccess1，该函数仅会返回一个指向目标值的指针；
+ 当接受两个参数时，会使用 runtime.mapaccess2，除了返回目标值之外，它还会返回一个用于表示当前键对应的值是否存在的 bool 值：

然后再根据KEY的不同类型,调用具体的 runtime.mapaccess1_fast64 , runtime.mapaccess1_fast32 , runtime.mapaccess1_faststr 等

```
func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	if raceenabled && h != nil {
		callerpc := getcallerpc()
		pc := funcPC(mapaccess1)
		racereadpc(unsafe.Pointer(h), callerpc, pc)
		raceReadObjectPC(t.key, key, callerpc, pc)
	}
	if msanenabled && h != nil {
		msanread(key, t.key.size)
	}
	if h == nil || h.count == 0 {
		if t.hashMightPanic() {
			t.hasher(key, 0) // see issue 23734
		}
		return unsafe.Pointer(&zeroVal[0])
	}
	if h.flags&hashWriting != 0 {
		throw("concurrent map read and map write")
	}
	hash := t.hasher(key, uintptr(h.hash0))
	m := bucketMask(h.B)
	b := (*bmap)(add(h.buckets, (hash&m)*uintptr(t.bucketsize)))
	if c := h.oldbuckets; c != nil {
		if !h.sameSizeGrow() {
			// There used to be half as many buckets; mask down one more power of two.
			m >>= 1
		}
		oldb := (*bmap)(add(c, (hash&m)*uintptr(t.bucketsize)))
		if !evacuated(oldb) {
			b = oldb
		}
	}
	top := tophash(hash)
bucketloop:
	for ; b != nil; b = b.overflow(t) {
		for i := uintptr(0); i < bucketCnt; i++ {
			if b.tophash[i] != top {
				if b.tophash[i] == emptyRest {
					break bucketloop
				}
				continue
			}
			k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
			if t.indirectkey() {
				k = *((*unsafe.Pointer)(k))
			}
			if t.key.equal(key, k) {
				e := add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.elemsize))
				if t.indirectelem() {
					e = *((*unsafe.Pointer)(e))
				}
				return e
			}
		}
	}
	return unsafe.Pointer(&zeroVal[0])
}
```

+ 第一段是竞态检测
+ 判断hmap是否为空,长度是否是0,若是返回零值
+ 判断当前是否并发读写map,若是则抛出异常
+ 如果hmap.B == 0 说明当前只有一个bucket,直接拿取.如果 > 0 , 就要计算当前key是落在哪个桶.首先对key + 哈希因子 进行hash运算,然后对桶的个
数进行取模运算,看落在哪个桶内.
+ 判断是否正在发生扩容,(判断h.oldbuckets是否为nil),如果正在扩容,就到老的buckets里面取查找
+ 遍历桶,对比tophash值(前八位),如果不等,continue, 如果元素等于0,直接break,说明后面空了.
+ tophash如果匹配成功,则正式完整的key进行比对,若查找成功返回.
+ 如果key是int类型 是直接用key进行比对,没必要用tophash先比对

## 写入
```
func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	if h == nil {
		panic(plainError("assignment to entry in nil map"))
	}
	if raceenabled {
		callerpc := getcallerpc()
		pc := funcPC(mapassign)
		racewritepc(unsafe.Pointer(h), callerpc, pc)
		raceReadObjectPC(t.key, key, callerpc, pc)
	}
	if msanenabled {
		msanread(key, t.key.size)
	}
	if h.flags&hashWriting != 0 {
		throw("concurrent map writes")
	}
	hash := t.hasher(key, uintptr(h.hash0))

	// Set hashWriting after calling t.hasher, since t.hasher may panic,
	// in which case we have not actually done a write.
	h.flags ^= hashWriting

	if h.buckets == nil {
		h.buckets = newobject(t.bucket) // newarray(t.bucket, 1)
	}

again:
	bucket := hash & bucketMask(h.B)
	if h.growing() {
		growWork(t, h, bucket)
	}
	b := (*bmap)(add(h.buckets, bucket*uintptr(t.bucketsize)))
	top := tophash(hash)

	var inserti *uint8
	var insertk unsafe.Pointer
	var elem unsafe.Pointer
bucketloop:
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			if b.tophash[i] != top {
				if isEmpty(b.tophash[i]) && inserti == nil {
					inserti = &b.tophash[i]
					insertk = add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
					elem = add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.elemsize))
				}
				if b.tophash[i] == emptyRest {
					break bucketloop
				}
				continue
			}
			k := add(unsafe.Pointer(b), dataOffset+i*uintptr(t.keysize))
			if t.indirectkey() {
				k = *((*unsafe.Pointer)(k))
			}
			if !t.key.equal(key, k) {
				continue
			}
			// already have a mapping for key. Update it.
			if t.needkeyupdate() {
				typedmemmove(t.key, k, key)
			}
			elem = add(unsafe.Pointer(b), dataOffset+bucketCnt*uintptr(t.keysize)+i*uintptr(t.elemsize))
			goto done
		}
		ovf := b.overflow(t)
		if ovf == nil {
			break
		}
		b = ovf
	}

	// Did not find mapping for key. Allocate new cell & add entry.

	// If we hit the max load factor or we have too many overflow buckets,
	// and we're not already in the middle of growing, start growing.
	if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
		hashGrow(t, h)
		goto again // Growing the table invalidates everything, so try again
	}

	if inserti == nil {
		// The current bucket and all the overflow buckets connected to it are full, allocate a new one.
		newb := h.newoverflow(t, b)
		inserti = &newb.tophash[0]
		insertk = add(unsafe.Pointer(newb), dataOffset)
		elem = add(insertk, bucketCnt*uintptr(t.keysize))
	}

	// store new key/elem at insert position
	if t.indirectkey() {
		kmem := newobject(t.key)
		*(*unsafe.Pointer)(insertk) = kmem
		insertk = kmem
	}
	if t.indirectelem() {
		vmem := newobject(t.elem)
		*(*unsafe.Pointer)(elem) = vmem
	}
	typedmemmove(t.key, insertk, key)
	*inserti = top
	h.count++

done:
	if h.flags&hashWriting == 0 {
		throw("concurrent map writes")
	}
	h.flags &^= hashWriting
	if t.indirectelem() {
		elem = *((*unsafe.Pointer)(elem))
	}
	return elem
}
```

+ 第一段是竞态检测
+ 判断当前是否并发读写map,若是则抛出异常,若不是则更改flag成写入状态
+ 根据传入的key拿到对应的哈希,哈希对桶个数取模,拿到对应的bucket,然后计算得到tophash
+ 扩容判断
+ 然后通过遍历比较桶中存储的 tophash 和键的哈希，如果找到了相同结果就会返回目标位置的地址。其中 inserti 表示目标元素的在桶中的索引，insertk
和 val 分别表示键值对的地址，获得目标地址之后会通过算术计算寻址获得键值对 k 和 val
+ 如果当前桶已经满了，哈希会调用 runtime.hmap.newoverflow 创建新桶或者使用 runtime.hmap 预先在 noverflow 中创建好的桶来保存数据，新创建
的桶不仅会被追加到已有桶的末尾，还会增加哈希表的 noverflow 计数器。

## 扩容
```
func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	...
	if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
		hashGrow(t, h)
		goto again
	}
	...
}
```
runtime.mapassign 函数会在以下两种情况发生时触发哈希的扩容：
1. 装载因子已经超过 6.5；翻倍扩容
2. 哈希使用了太多溢出桶；等量扩容
   1. B < 15 时，溢出桶数量 >= 2^B
   2. B >= 15 时，溢出桶数量 >= 2^15
触发等量扩容，是在装载因子较小的情况下， Map 的读写效率也较低。这种现象是 Map 元素少，但溢出桶数量很多，等量扩容后，分配的新桶数量跟旧桶数量相同，但新桶中存储的数据更加紧密

   3. 不过因为 Go 语言哈希的扩容不是一个原子的过程，所以 runtime.mapassign 还需要判断当前哈希是否已经处于扩容状态，避免二次扩容造成混乱。

根据触发的条件不同扩容的方式分成两种，如果这次扩容是溢出的桶太多导致的，那么这次扩容就是等量扩容 sameSizeGrow，sameSizeGrow 是一种特殊情况下
发生的扩容，当我们持续向哈希中插入数据并将它们全部删除时，如果哈希表中的数据量没有超过阈值，就会不断积累溢出桶造成缓慢的内存泄漏, 
引入了 sameSizeGrow 通过复用已有的哈希扩容机制解决该问题，一旦哈希中出现了过多的溢出桶，它会创建新桶保存数据，垃圾回收会清理老的溢出桶并释放内存

```
func hashGrow(t *maptype, h *hmap) {
	bigger := uint8(1)
	if !overLoadFactor(h.count+1, h.B) {
		bigger = 0
		h.flags |= sameSizeGrow
	}
	oldbuckets := h.buckets
	newbuckets, nextOverflow := makeBucketArray(t, h.B+bigger, nil)

	h.B += bigger
	h.flags = flags
	h.oldbuckets = oldbuckets
	h.buckets = newbuckets
	h.nevacuate = 0
	h.noverflow = 0

	h.extra.oldoverflow = h.extra.overflow
	h.extra.overflow = nil
	h.extra.nextOverflow = nextOverflow
}
```

哈希在扩容的过程中会通过 runtime.makeBucketArray 创建一组新桶和预创建的溢出桶，随后将原有的桶数组设置到 oldbuckets 上并将新的空桶设置到 
buckets 上，溢出桶也使用了相同的逻辑更新

我们在 runtime.hashGrow 中还看不出来等量扩容和翻倍扩容的太多区别，等量扩容创建的新桶数量只是和旧桶一样，该函数中只是创建了新的桶，并没有对数据
进行拷贝和转移。哈希表的数据迁移的过程在是 runtime.evacuate 中完成的，它会对传入桶中的元素进行再分配
```
func evacuate(t *maptype, h *hmap, oldbucket uintptr) {
	b := (*bmap)(add(h.oldbuckets, oldbucket*uintptr(t.bucketsize)))
	newbit := h.noldbuckets()
	if !evacuated(b) {
		var xy [2]evacDst
		x := &xy[0]
		x.b = (*bmap)(add(h.buckets, oldbucket*uintptr(t.bucketsize)))
		x.k = add(unsafe.Pointer(x.b), dataOffset)
		x.v = add(x.k, bucketCnt*uintptr(t.keysize))

		y := &xy[1]
		y.b = (*bmap)(add(h.buckets, (oldbucket+newbit)*uintptr(t.bucketsize)))
		y.k = add(unsafe.Pointer(y.b), dataOffset)
		y.v = add(y.k, bucketCnt*uintptr(t.keysize))
```
runtime.evacuate 会将一个旧桶中的数据分流到两个新桶，所以它会创建两个用于保存分配上下文的 runtime.evacDst 结构体，这两个结构体分别指向了一个新桶

## 删除
哈希表的删除逻辑与写入逻辑很相似，只是触发哈希的删除需要使用关键字，如果在删除期间遇到了哈希表的扩容，就会分流桶中的元素，分流结束之后会找到桶中的目标元素完成键值对的删除工作

