package sort

import (
	"math/rand"
	"testing"
	"time"
)

// 利用 ^ 运算,实现不分配额外内存空间交换两个值
func TestArraySwap(t *testing.T) {
	i := 1
	j := 2

	i = i ^ j
	j = i ^ j // j = i ^ j ^ j  ==> j = i
	i = i ^ j // i = i ^ j ^ i ==> j = j

	t.Log(i, j)
}

// 生成一个随机数组
func randSlice(n int) (s []int) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(n))
	}

	return
}

// 选择排序
// 时间复杂度 0(n^2)
//名字由来: 第一次从待排序的数据元素中选出最小（或最大）的一个元素，存放在序列的起始位置，然后再从剩余的未排序元素中寻找到最小（大）元素，
//然后放到已排序的序列的末尾。以此类推，直到全部待排序的数据元素的个数为零。选择排序是不稳定的排序方法。
func SelectionSort(s []int) {
	for i := 0; i < len(s); i++ {
		// 等差数列,每次选择最小放在范围队列起始位置
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

func TestSelectionSort(t *testing.T) {
	s := randSlice(200)
	SelectionSort(s)
	t.Log(s)
}
