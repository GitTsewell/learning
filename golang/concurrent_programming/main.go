package main

import "fmt"

func main() {
	var b uint64
	for {
		a := 1
		b = uint64(a) + 1
	}
	fmt.Println(b)
}
