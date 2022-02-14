package main

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	ch := make(chan int)
	ch <- 3
}

func TestNilChannel(t *testing.T) {
	var ch chan int

	go func() {
		fmt.Println(<-ch)
	}()

	time.Sleep(1 * time.Second)

	ch <- 1

	time.Sleep(1 * time.Second)
}
