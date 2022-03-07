package slice

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	// 使用各自独立的 6 个 slice 来创建 [2][3] 的动态多维数组
	x := 2
	y := 4

	table := make([][]int, x)
	for i := range table {
		table[i] = make([]int, y)
	}

	fmt.Println(table)
}

func TestMulArray(t *testing.T) {
	arr := [2][2]int{{2, 2}, {2, 2}}
	fmt.Println(arr)
}
