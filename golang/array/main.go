package main

import "fmt"

func Assign1(s []int) {
	s = []int{6, 6, 6}
}

func Reverse0(s [5]int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse1(s []int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse2(s []int) {
	s = append(s, 999)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse3(s []int) {
	s = append(s, 999, 1000, 1001)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	Assign1(s)
	fmt.Println(s) // (1)

	array := [5]int{1, 2, 3, 4, 5}
	Reverse0(array)
	fmt.Println(array) // (2)

	s = []int{1, 2, 3}
	Reverse2(s)
	fmt.Println(s) // (3)

	var a []int
	for i := 1; i <= 3; i++ {
		a = append(a, i)
	}
	Reverse2(a)
	fmt.Println(a) // (4)

	var b []int
	for i := 1; i <= 3; i++ {
		b = append(b, i)
	}
	Reverse3(b)
	fmt.Println(b) // (5)

	c := [3]int{1, 2, 3}
	d := c
	c[0] = 999
	fmt.Println(d) // (6)
}
