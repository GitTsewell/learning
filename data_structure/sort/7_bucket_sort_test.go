package sort

import (
	"fmt"
	"testing"
)

// 桶排序
// 前提条件是在一个适当的范围,不急于比较的排序
// 时间复杂度 O(n+k) 空间复杂度 O(k)  稳定
// 比如[4,4,4,6,7,7,7,4,3,3,3,2]这样一组序列,范围都在[0~9]之间,遍历一次分别记下有多少个1,多少个2,多少个3,最后遍历出来就有序了
func BucketSort(s []int, l, r int) []int {
	// 初始化一个map k是数字,v是出现的次数
	m := make(map[int]int, r-l)

	for i := l; i <= r; i++ {
		m[i] = 0
	}

	// 遍历arr 记录出现次数
	for _, v := range s {
		m[v] = m[v] + 1
	}

	// 因为golang map是无序的 所以不能range 依赖l-->r 顺序遍历
	var newS []int
	for i := l; i <= r; i++ {
		if _, ok := m[i]; ok && m[i] != 0 {
			for j := 0; j < m[i]; j++ {
				newS = append(newS, i)
			}
		}
	}

	return newS
}

func TestBucketSort(t *testing.T) {
	s := []int{4, 5, 8, 8, 4, 3, 6, 5, 4, 3, 6, 6, 5, 3, 7, 8, 5, 4, 3, 2, 4, 6, 3}
	s = BucketSort(s, 0, 9)
	fmt.Println(s)
}
