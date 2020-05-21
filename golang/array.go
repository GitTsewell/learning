package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", &arr[1])
	fmt.Printf("%p\n", &arr[2])
}
