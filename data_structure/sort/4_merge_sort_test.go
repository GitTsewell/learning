package sort

import (
	"learning/data_structure"
	"testing"
)

//归并排序（英语：Merge sort，或mergesort），是创建在归并操作上的一种有效的排序算法。1945年由约翰·冯·诺伊曼首次提出。
//该算法是采用分治法（Divide and Conquer）的一个非常典型的应用，且各层分治递归可以同时进行。
// 时间复杂度 O(nlog n)  空间复杂度 O(n)
// 分两个维度,第一个归,也就是递归,先把序列平均分成两半,直到这个序列个数只有一个,
// 然后调用merge合并,合并的具体步骤如下
// 比如a[1,3,5]和b[2,4,6]合并,那先从a[0],b[0]比较大小合并,a[0]比b[1]小,那a[0]先放到一个新序列n中,然后a[1]和b[0]在比较大小,直到序列比较完

func merge(a, b []int) []int {
	// 预先分配好内存空间,避免append之后的切片扩容
	s := make([]int, len(a)+len(b))

	i := 0
	j := 0
	for {
		// a,b都越界 说明完成了  直接break
		if i > len(a)-1 && j > len(b)-1 {
			break
		}
		// a的越界判断,如果i >= len(a)-1,那么把b剩下的序列依次追加进s中
		if i > len(a)-1 {
			s[i+j] = b[j]
			j++
			continue
		}
		// 同理b的越界判断
		if j > len(b)-1 {
			s[i+j] = a[i]
			i++
			continue
		}
		// 比较a b 当前offset的大小,如果a[i] <= b[j],那么先把a[i]放到s[i+j]中
		if a[i] <= b[j] {
			s[i+j] = a[i]
			i++
		} else {
			s[i+j] = b[j]
			j++
		}
	}

	return s
}

func mergeSort(s []int) []int {
	// 找到mid 一分为二,然后merge,所有的一分为二的子集,再merge,直到 len(s) < 2,直接返回
	if len(s) < 2 {
		return s
	}
	mid := len(s) >> 1
	return merge(mergeSort(s[:mid]), mergeSort(s[mid:]))
}

func TestMerge(t *testing.T) {
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8, 9}
	s := merge(a, b)
	t.Log(s)
}

func TestMergeSort(t *testing.T) {
	s := data_structure.RandSlice(200)
	s = mergeSort(s)
	t.Log(s)
}
