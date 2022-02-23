package binary_search

import (
	"learning/data_structure"
	"sort"
	"testing"
)

// 二分查找算法
// 又叫折半搜索算法,或对数搜索算法.是一种在有序数组中查找某一特定元素的搜索算法。搜索过程从数组的中间元素开始，如果中间元素正好是要查找的元素，
//则搜索过程结束；如果某一特定元素大于或者小于中间元素，则在数组大于或小于中间元素的那一半中查找，而且跟开始一样从中间元素开始比较。如果在某一步骤数组为空，
//则代表找不到。这种搜索算法每一次比较都使搜索范围缩小一半

// 递归 终止条件: 1.找到这个数  2.左边界 > 右边界
func binarySearchRecursion(s []int, target, l, r int) int {
	if l > r {
		return -1
	}

	// 注意不要用 mid=(l+r)/2 这种方式,容易造成溢出
	mid := l + (r-l)/2
	// 如果中间值大于target,就往左边界查找
	if s[mid] > target {
		return binarySearchRecursion(s, target, l, mid-1)
	} else if s[mid] < target {
		return binarySearchRecursion(s, target, mid+1, r)
	}

	return mid
}

func TestBinarySearchRecursion(t *testing.T) {
	s := data_structure.RandSlice(200)
	sort.Ints(s)
	target := 101
	offset := binarySearchRecursion(s, target, 0, 199)
	t.Log(s)
	if offset != -1 && s[offset] == target {
		t.Logf("数组里面包含目标值,offset是 : %d", offset)
	} else {
		t.Log("目标数组里面不包含目标值")
	}
}

func binarySearchNotRecursion(s []int, target int) int {
	if len(s) == 0 {
		return -1
	}

	l := 0
	r := len(s) - 1

	for r > l {
		// 求出mid 比较s[mid] 和 target 大小 确定左右边界
		mid := l + (r-l)/2
		if s[mid] == target {
			return mid
		} else if s[mid] > target {
			r = mid - 1
			continue
		} else {
			l = mid + 1
			continue
		}
	}

	return -1
}

func TestBinarySearchNotRecursion(t *testing.T) {
	s := data_structure.RandSlice(200)
	sort.Ints(s)
	target := 101
	offset := binarySearchNotRecursion(s, target)
	t.Log(s)
	if offset != -1 && s[offset] == target {
		t.Logf("数组里面包含目标值,offset是 : %d", offset)
	} else {
		t.Log("目标数组里面不包含目标值")
	}
}
