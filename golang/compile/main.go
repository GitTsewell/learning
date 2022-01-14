package main

const name uint64 = 100

func main() {
	//当元素数量小于或者等于 4 个时，会直接将数组中的元素放置在栈上；
	//当元素数量大于 4 个时，会将数组中的元素放置到静态区并在运行时取出；
	a := [4]interface{}{}
	b := [100]interface{}{}

	_ = a
	_ = b
}
