package array

import (
	"fmt"
	"testing"
)

// 二维切片的内存地址
func TestMulArray(t *testing.T) {
	type arr [3]int
	var array [3]arr
	array[0] = arr{1, 2, 3}
	array[1] = arr{1, 2, 3}
	array[2] = arr{1, 2, 3}
	fmt.Printf("%p\n", &array)
	fmt.Printf("%p\n", &array[0][0])
	fmt.Printf("%p\n", &array[0][1])
	fmt.Printf("%p\n", &array[0][2])
	fmt.Printf("%p\n", &array[1])
	fmt.Printf("%p\n", &array[2])
}
