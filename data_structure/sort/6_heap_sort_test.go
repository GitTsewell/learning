package sort

import (
	"learning/data_structure"
	"testing"
)

// 堆排序,是指利用堆这种数据结构来进行排序的手段
// 堆是一种完全二叉树,又分为大根堆和小根堆
// 堆中几个概念: heapSize 堆的大小  i的左节点 = 2i+1 , i的右节点 = 2i + 2,  i的父节点 = (i-1)/2
// 在堆的操作过程中,主要用两种方法.  heapInsert  和 heapify
// heapInset操作 : 比如向一个大根堆插入一个新数据X,heapSize+1 寻找父节点PK,如果比父节点大就swap,直到父节点>=自己就停止
// heapify操作 : 移除大根堆的定点,然后重排序,让这个结构重新成为一个大根堆.
// 首先把最后一个数和定点swap,heapSize-1,然后判断是否有子节点(可以先判断左节点,如果左节点都没有,那肯定没有右节点),如果有左节点,那么判断有没有右节点,
// 如果有右节点就左右节点先比较,选出较大节点和该节点比较,如果该节点大,就stop,如果左(右)节点大,就swap,然后再判断上述过程,知道没有子节点或子节点都比自己小

// // heapInset操作 : 比如向一个大根堆插入一个新数据X,heapSize+1 寻找父节点PK,如果比父节点大就swap,直到父节点>=自己就停止
func heapInsert(s []int, num int) []int {
	s = append(s, num)
	heapSize := len(s)

	var i = heapSize - 1
	for i > 0 {
		pi := (i - 1) >> 1
		if s[pi] >= s[i] {
			break
		} else {
			s[pi], s[i] = s[i], s[pi]
			i = pi
		}
	}
	return s
}

func TestHeapInsert(t *testing.T) {
	s := []int{6, 5, 4, 3, 2, 1}
	num := 7
	s = heapInsert(s, num)
	t.Log(s)
}

// heapify操作 : 移除大根堆的定点,然后重排序,让这个结构重新成为一个大根堆.
// 首先把最后一个数和定点swap,heapSize-1,然后判断是否有子节点(可以先判断左节点,如果左节点都没有,那肯定没有右节点),如果有左节点,那么判断有没有右节点,
// 如果有右节点就左右节点先比较,选出较大节点和该节点比较,如果该节点大,就stop,如果左(右)节点大,就swap,然后再判断上述过程,知道没有子节点或子节点都比自己小
func heapify(s []int, heapSize int) int {
	endIndex := heapSize - 1
	s[0], s[endIndex] = s[endIndex], s[0]
	heapSize--

	i := 0
	for (2*i + 1) < heapSize {
		index := 2*i + 1
		// 判断是否有右节点,并且右节点比左节点大,就取右节点
		if (2*i+2) < heapSize && s[2*i+2] > s[index] {
			index = 2*i + 2
		}

		// 判断i节点大 直接break
		if s[i] >= s[index] {
			break
		} else {
			s[i], s[index] = s[index], s[i]
			i = index
		}
	}
	return heapSize
}

func TestHeapify(t *testing.T) {
	s := []int{7, 6, 5, 4, 3, 2, 1}
	heapify(s, len(s))
	t.Log(s)
}

// 堆排序
// 时间复杂度 O(n*logN) 空间复杂度 O(1) 不稳定
// arr[0]天然是一个大根堆,然后一次插入arr[i] [i+1] 一直到i = len(arr)-1,就相当于循环调用heapInsert过程  这里时间复杂度 O(n*log n)
// 然后取出根顶,和arr[heapSize]交换,就相当于取出最大值,放在数组最后,然后heapSize-1,因为末尾已经是最大值了,不用维护了,然后循环这个过程,直到heapSize等于0,
// 这里就相当于做 heapify操作,时间复杂度 O(n*log n)
func HeapSort(s []int) []int {
	if len(s) < 2 {
		return s
	}

	n := []int{s[0]}
	for i := 1; i < len(s); i++ {
		n = heapInsert(n[:i], s[i])
	}

	heapSize := len(n)
	for heapSize > 0 {
		heapSize = heapify(n, heapSize)
	}
	return n
}

func TestHeapSort(t *testing.T) {
	s := data_structure.RandSlice(20)
	t.Log(s)
	s = HeapSort(s)
	t.Log(s)
}

// 堆排序扩展题
// 已知一个几乎有序的数组,几乎有序是指,如果把数组排序号,每一个元素的移动步数不会超过k,并且k相对于数组很小,请排序
// 最大移动不会超过k,所以数组的最小值肯定在[0,k]中,所以现在arr[0,k]中,arr[0,k]变成一个小根堆,然后把堆顶放在arr[0]位置,然后重复,直到i=len(arr)-1

func heapInsertMin(s []int, num int) []int {
	s = append(s, num)
	heapSize := len(s)

	var i = heapSize - 1
	for i > 0 {
		pi := (i - 1) >> 1
		if s[pi] <= s[i] {
			break
		} else {
			s[pi], s[i] = s[i], s[pi]
			i = pi
		}
	}
	return s
}

func HeapSortQuestionK(s []int, k int) []int {
	var newS []int
	for i := 0; i < len(s); i++ { // 确定lIndex
		// 确定k
		if len(s[i:]) < k {
			k = len(s[i:])
		}

		ranges := s[i : i+k]
		n := []int{s[i]}
		for j := 1; j < len(ranges); j++ {
			n = heapInsertMin(ranges[:j], ranges[j])
		}
		newS = append(newS, n[0])
	}
	return newS
}

func TestHeapSortQuestionK(t *testing.T) {
	s := []int{3, 2, 4, 5, 8, 6, 9, 7}
	s = HeapSortQuestionK(s, 4)
	t.Log(s)
}
