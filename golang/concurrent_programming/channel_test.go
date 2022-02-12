package main

import "testing"

func Test(t *testing.T) {
	ch := make(chan int)
	ch <- 3
}
