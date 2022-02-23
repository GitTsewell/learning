package sort

import "testing"

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
