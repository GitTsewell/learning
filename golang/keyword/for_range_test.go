package keyword

import (
	"fmt"
	"testing"
)

func TestSplitRange(t *testing.T) {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
}

func TestMapRange(t *testing.T) {
	m := map[int]int{
		2: 2,
	}
	for _, v := range m {
		j := v << 1
		m[j] = j
	}
	fmt.Println(m)
}
