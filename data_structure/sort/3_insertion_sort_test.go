package sort

import "testing"

// 插入排序
// 时间复杂度 O(n^2) 空间复杂度 0(1) 稳定
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
