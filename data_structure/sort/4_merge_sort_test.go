package sort

import (
	"learning/data_structure"
	"testing"
)

// 归并排序（英语：Merge sort，或mergesort），是创建在归并操作上的一种有效的排序算法。1945年由约翰·冯·诺伊曼首次提出。
// 核心思想:分而治之,先让其局部有序,然后合并局部直到整个序列
// 该算法是采用分治法（Divide and Conquer）的一个非常典型的应用，且各层分治递归可以同时进行。
// 时间复杂度 O(nlog n)  空间复杂度 O(n)
// 分两个维度,第一个归,也就是递归,先把序列平均分成两半,直到这个序列个数只有一个,
// 然后调用merge合并,合并的具体步骤如下
// 比如a[1,3,5]和b[2,4,6]合并,那先从a[0],b[0]比较大小合并,a[0]比b[1]小,那a[0]先放到一个新序列n中,然后a[1]和b[0]在比较大小,直到序列比较完

// 整体就是一个简单的递归,左边排好序,右边排好序,然后让其整体有序
// 让其整体有序的过程采用了外排序
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

// 归并排序扩展的一些问题
// 小和问题 : 在一个数组中,每一个数左边比当前数小的数累加起来,叫做这个数组的小和.求一个数组的小和
// 例子:[1,3,4,2,5],比1小的没有,比3小的1 比4小的1,3 比2小的1,比5小的1,3,4,2  加起来1+1+3+1+1+3+4+2 = 16

// 自己思路
// 可以采用双重遍历的暴力算法,依次去比较,然后得出小和 时间复杂度O(n^2)
// 可以采用归并排序,为什么? 首先换个角度考虑,看一个数右边有几个数比他大,这个数在小和中就会出来几次,比如1右边有四个数都比1大,那就是4个1,3,右边有两个数比3大就是2个3,依次类推
// 然后利用归并局部有序的特性,可以少比较很多次,比如a[0]和b[0]比较,如果a[0]<b[0] 那么就会有b序列中所有数都大于a[0],所以a[0]在这个局部出现次数是len(0)个,以此类推

var sum int

func mergeMinSum(a, b []int) []int {
	// 分配内存空间
	s := make([]int, len(b)+len(a))

	i := 0
	j := 0

	for {
		if i > len(a)-1 && j > len(b)-1 {
			break
		}

		if i > len(a)-1 {
			s[i+j] = b[j]
			j++
			continue
		}
		if j > len(b)-1 {
			s[i+j] = a[i]
			i++
			continue
		}

		if a[i] < b[j] {
			s[i+j] = a[i]
			sum += a[i] * (len(b) - j)
			i++
		} else {
			// 左边数比右边大,说明局部右边b[j]都比左边小 所以sum不用加 跳过j 等下一次j+1 在判断
			s[i+j] = b[j]
			j++
		}
	}
	return s
}

func TestMergeMinSum(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}

	s := mergeMinSum(a, b)
	t.Log(s, sum)
}

func MergeSortMinSum(s []int) []int {
	if len(s) < 2 {
		return s
	}

	mid := len(s) >> 1

	return mergeMinSum(MergeSortMinSum(s[:mid]), MergeSortMinSum(s[mid:]))
}

func TestMergeSortMinSum(t *testing.T) {
	s := data_structure.RandSlice(4)
	t.Log(s)
	s = MergeSortMinSum(s)
	t.Log(s)
	t.Log(sum)
}
