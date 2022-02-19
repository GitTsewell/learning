package main

import (
	"fmt"
	"sync"
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

type Student struct {
	name string
}

func TestMapValue(t *testing.T) {
	m := map[string]Student{"people": {"zhoujielun"}}
	m["people2"] = Student{name: "pp"}
	c := m["people"]
	fmt.Println(c)
}

func TestSyncMap(t *testing.T) {
	m := sync.Map{}
	m.Store(1, 1)
	m.Store(2, 2)
	m.Load(1)
	m.Load(2)
	m.Store(3, 3)
	m.Store(4, 4)
	m.Load(3)
}

func TestMapAdress(t *testing.T) {
	m1 := map[int]int{1: 1}
	m2 := m1

	fmt.Printf("m1 address : %p,m2 address : %p\n", m1, m2)

	for i := 10; i <= 10000; i++ {
		m1[i] = i
	}

	fmt.Printf("m1 address : %p,m2 address : %p\n", m1, m2)
}
