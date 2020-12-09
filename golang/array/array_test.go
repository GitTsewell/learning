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

func TestCapSlice(t *testing.T) {
	sli := make([]int, 0, 0)
	t.Logf("sli len is %d , cap is %d , address is %p \n", len(sli), cap(sli), &sli)

	sli = append(sli, 1)
	t.Logf("after append 1 sli len is %d , cap is %d , address is %p \n", len(sli), cap(sli), &sli)

	sli = append(sli, 2, 3)
	t.Logf("after append 2 sli len is %d , cap is %d , address is %p \n", len(sli), cap(sli), &sli)

	sli = append(sli, 4, 5, 6)
	t.Logf("after append 3 sli len is %d , cap is %d , address is %p \n", len(sli), cap(sli), &sli)

	sli = append(sli, 7, 8, 9, 10)
	t.Logf("after append 4 sli len is %d , cap is %d , address is %p \n", len(sli), cap(sli), &sli)

	// the array cap is 0
	var arr = [0]int{}
	t.Logf("array len is %d , cap is %d , address is %p \n", len(arr), cap(arr), &arr)
	//var ptr unsafe.Pointer
	//var o []byte
	//sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
	//sliceHeader.Cap = 0
	//sliceHeader.Len = 0
	//sliceHeader.Data = uintptr(ptr)
}

func TestA(t *testing.T) {
	var a int8 = 1
	t.Log(*(*int8)(&a))
}
