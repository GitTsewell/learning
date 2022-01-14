package main

import "fmt"

func main() {
	m := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}

	m1 := m["1"]
	fmt.Println(m1)
}
