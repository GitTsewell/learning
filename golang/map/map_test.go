package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMapRead(t *testing.T) {
	m := map[int]int{
		1: 1,
	}

	for i := 2; i < 1000; i++ {
		m[i] = i
	}
	
	m1 := m[1]
	fmt.Println(m1)
}

func TestName(t *testing.T) {
	k := 1
	fmt.Println(unsafe.Pointer(&k))
}
