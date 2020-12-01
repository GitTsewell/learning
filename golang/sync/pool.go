package main

import (
	"fmt"
	"sync"
)

type Pool struct {
	Pool *sync.Pool
}

// 临时对象池,减少创建临时变量
func main() {
	ts := &Pool{Pool: &sync.Pool{New: func() interface{} {
		return 1
	}}}
	ts.Put(1)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			r := ts.Get()
			fmt.Println(r)
			wg.Done()
		}()
	}
	wg.Wait()
}

func (slf *Pool) Get() interface{} {
	r := slf.Pool.Get()
	if r != nil {
		defer slf.Pool.Put(r.(int) + 1)
	}
	return r
}

func (slf *Pool) Put(x interface{}) {
	slf.Pool.Put(x)
}
