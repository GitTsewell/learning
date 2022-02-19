package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestNewArray(t *testing.T) {
	a := [20]int{}
	fmt.Println(a)
}

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

	sli = append(sli, 11, 12, 13, 14, 15)
	t.Logf("after append 4 sli len is %d , cap is %d , address is %p \n", len(sli), cap(sli), &sli)

	sli = sli[0:1]
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
	sli1 := make([]int, 1024)
	sli1 = append(sli1, 1024, 1025)
	t.Logf("after append 4 sli len is %d , cap is %d , address is %p , size is %v \n", len(sli1), cap(sli1), &sli1, unsafe.Sizeof(sli1))

	arr1 := [1024]int{}
	t.Logf("size is %v \n", unsafe.Sizeof(arr1))
}

func TestArrayElementLength(t *testing.T) {
	a := [200]interface{}{}
	fmt.Println(&a[0], &a[1])

	b := [200]string{}
	fmt.Println(&b[0], &b[1])

	type s struct {
		X string      `json:"x"`
		Y uint64      `json:"y"`
		Z interface{} `json:"z"`
	}

	c := [200]s{}
	fmt.Printf("%p and %p \n", &c[0], &c[1])
}

func TestArrayEq(t *testing.T) {
	a := [2]int{5, 6}
	b := [2]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

func TestSliceCopy(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := a
	b[0] = 100

	fmt.Println(a, b)
	fmt.Printf("a addres is %p , b address is %p \n", &a[0], &b[0])

	c := []int{1, 2, 3, 4}
	d := make([]int, 4)
	copy(d, c)
	d[0] = 100
	fmt.Println(c, d)
	fmt.Printf("c addres is %p , d address is %p \n", &c[0], &d[0])
}

func TestSliceAdd(t *testing.T) {
	a := []int{0}
	b := a

	fmt.Printf("a address : %p,b address : %p\n", a, b)

	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)

	fmt.Printf("a address : %p,b address : %p\n", a, b)

}
