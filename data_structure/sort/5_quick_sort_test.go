package sort

import (
	"learning/data_structure"
	"testing"
)

// 在快速排序之前,先解决两个小问题
// 1.给定一个数组和一个数num,把小于等于num的放左边,大于num的放右边,要求时间复杂度O(n)
// 解决方案: 1;暴力算法,全部循环一遍
// 单指针: 指针一开始指向[0],然后遍历数组,只要有一个元素小于等于num就把当前i位和指针指向那一位交换,指针往前移动一位 i++, 如果大于就不移动指针,只是i++
func singlePoint(s []int, target int) {
	p := 0
	for i, _ := range s {
		if s[i] <= target {
			s[p], s[i] = s[i], s[p]
			p++
		}
	}
}

func TestSinglePoint(t *testing.T) {
	s := data_structure.RandSlice(20)
	singlePoint(s, 8)
	t.Log(s)
}

// 第二个问题
// 给定一个数组arr 和一个数num,请把小于num放左边,等于num的放中间,大于num的放右边,要求时间复杂度O(n)
// 这种一个整体分成三部分的,可以考虑使用双指针
// lp指向[0],rp指向序列尾,如果arr[i]<num,arr[i] swap arr[lp],然后lp+1 i+1,如果arr[1]==num,不用动 i+1,如果arr[i]>num,arr[i] swap arr[rp},rp-1
// i 不加1,因为换过来的数是还没验证过得,最后lp左边的都是小于num lp和rp中间的都是等于num,rp右边的都是大于num的,当i==rp 停止
func doublePoint(s []int, target int) {
	if len(s) < 2 {
		return
	}

	lp := 0
	rp := len(s) - 1

	for i := 0; i < len(s); {
		if i >= rp {
			break
		}
		if s[i] < target {
			s[lp], s[i] = s[i], s[lp]
			lp++
			i++
		} else if s[i] == target {
			i++
		} else {
			s[rp], s[i] = s[i], s[rp]
			rp--
		}
	}
}

func TestDoublePoint(t *testing.T) {
	s := data_structure.RandSlice(20)
	t.Log(s)
	doublePoint(s, 6)
	t.Log(s)
}

// 快速排序
// 时间复杂度 O(n*log n)  最坏情况O(n^2)  空间复杂度O(logN) 不稳定
// 步骤,选一个数(数组最后一个)做基准,然后小于等于这个数的放左边,大于这个数的放右边,然后继续递归左边和右边的直到剧本序列<=2
// 步骤,数组最后一个数做target,arr[i]<=target放point左边 然后point++ i++,arr[i]>point 放右边 i++,最后吧最后一个数和point数做交换,
// 然后把 小于point的 和 大于point的再继续
func quickSort(s []int, l, r int) {
	if l == r {
		return
	}

	target := s[r]
	p := l
	for i := l; i < r; i++ {
		if s[i] <= target {
			s[i], s[p] = s[p], s[i]
			p++
		}
	}

	s[p], s[r] = s[r], s[p]

	if p > l {
		quickSort(s, l, p-1)
	}

	if p < r {
		quickSort(s, p+1, r)
	}
}

func TestQuickSort(t *testing.T) {
	s := data_structure.RandSlice(17)
	t.Log(s)
	quickSort(s, 0, len(s)-1)
	t.Log(s)
}

// 快速排序2.0
// 把序列分成三部分,小于num的放左边,等于num的放中间,大于num的放右边,两个指针,一个指向左边初始,一个指向右边初始
// 具体步骤 arr[i] < num swap i++ lp++ , arr[i] == num i++ , arr[i] > num swap rp--
func QuickSortV2(s []int, lIndex, rIndex int) {
	if lIndex >= rIndex {
		return
	}

	target := s[rIndex]
	lp := lIndex
	rp := rIndex

	for i := lIndex; lp < rp; {
		if s[i] < target {
			s[i], s[lp] = s[lp], s[i]
			lp++
			i++
		} else if s[i] == target {
			i++
		} else {
			s[i], s[rp] = s[rp], s[i]
			rp--
		}
	}

	if lp > lIndex {
		quickSort(s, lIndex, lp-1)
	}

	if rp < rIndex {
		quickSort(s, lp+1, rIndex)
	}
}

func TestQuickSortV2(t *testing.T) {
	s := data_structure.RandSlice(17)
	t.Log(s)
	QuickSortV2(s, 0, len(s)-1)
	t.Log(s)
}
