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

// 冒泡排序  又名泡式排序
// 时间复杂度 0(n^2)
// 这个算法的名字由来是因为越小的元素会经由交换慢慢“浮”到数列的顶端。
// 在范围内比较相邻的两个值大小,小的在前,大的在后,第一轮扫完之后,最大的会被排到最后,然后排除最后一位,缩小扫描范围,知道遍历完成,
//冒泡相对于选择会让前面已经扫描过得相对有序,某些情况下会减少后面扫描的swap次数
func BubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 1; j < len(s)-i; j++ {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
}

func TestBubbleSort(t *testing.T) {
	s := randSlice(200)
	BubbleSort(s)
	t.Log(s)
}

// 插入排序
// 它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
// 范围由小到大,先比较[0]范围,0范围[0]先天最小,然后加入[1] 和之前已经排序好的[0]比较,选择插入位置,是[0]之前还是之后,然后加入[2],比较之前
// 已经排序好的[0,1],从后往前比较,选择合适的地方插入
func InsertionSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if s[j] >= s[j-1] {
				continue
			} else {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
}

func TestInsertionSort(t *testing.T) {
	s := randSlice(200)
	InsertionSort(s)
	t.Log(s)
}
