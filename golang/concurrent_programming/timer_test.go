package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)
	//timer.Stop()
	for {
		//timer.Reset(3 * time.Second) // 这里复用了 timer
		select {
		case <-timer.C:
			fmt.Println("3秒后执行一次")
		}
	}
}

func TestGC(t *testing.T) {
	for {
		a := 1
		_ = a
	}
}
