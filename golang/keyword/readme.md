# for AND range
经典的for循环和for range 在编译后都是 ```for init;left;right {body}``` 的结构,没有区别
## 三个现象
### 循环永动机
如果我们在遍历数组的同时修改数组的元素，能否得到一个永远都不会停止的循环呢
```
func main() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
}

$ go run main.go
1 2 3 1 2 3
```
对于所有的 range 循环，Go 语言都会在编译期将原切片或者数组赋值给一个新变量 ha，在赋值的过程中就发生了拷贝，而我们又通过 len 关键字预先获取了切片的长度，
所以在循环中追加新的元素也不会改变循环执行的次数，这也就解释了循环永动机一节提到的现象

### 神奇的指针
第二个例子是使用 Go 语言经常会犯的错误1。当我们在遍历一个数组时，如果获取 range 返回变量的地址并保存到另一个数组或者哈希时，会遇到令人困惑的现象，下面的代码会输出 “3 3 3”：
```
func main() {
	arr := []int{1, 2, 3}
	newArr := []*int{}
	for _, v := range arr {
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}

$ go run main.go
3 3 3
```
而遇到这种同时遍历索引和元素的 range 循环时，Go 语言会额外创建一个新的 v2 变量存储切片中的元素，循环中使用的这个变量 v2 会在每一次迭代被重新赋值而覆盖，赋值时也会触发拷贝
因为在循环中获取返回变量的地址都完全相同，所以会发生神奇的指针一节中的现象。因此当我们想要访问数组中元素所在的地址时，不应该直接获取 range 返回的变量地址 &v2，而应该使用 &a[index] 这种形式

### map随机遍历
当我们在 Go 语言中使用 range 遍历哈希表时，往往都会使用如下的代码结构，但是这段代码在每次运行时都会打印出不同的结果：
```
func main() {
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range hash {
		println(k, v)
	}
}
```
Go 团队在设计哈希表的遍历时就不想让使用者依赖固定的遍历顺序，所以引入了随机数保证遍历的随机性
所以在循环开始的时候会生成一个随机数,确定bucket的位置,然后再通过这个随机数确定offset位置,最后bucket遍历完成之后再遍历overflow溢出桶

# select
1. 非租塞的收发
2. 两个case命中,随机执行