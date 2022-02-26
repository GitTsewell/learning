package sort

import (
	"fmt"
	"testing"
)

// 基数排序
// 时间复杂度 O(n+k) 空间复杂度 O(k)  稳定
// 比如一个序列都是0~99以内的数,先创建个位数字的十个桶,[0~9],每个桶里面是一个队列,如果一个数个位是1 就落在[1]号桶入队列,重复遍历完,然后
// 按照[0~9的顺序]一次把桶内的数倒出来,这样我们就完成的个位数的排序,然后再进行十位数的排序....
// 原理就是按照权重 权重最小的各位先排序

// 先实现一个队列

type queue struct {
	S []int
}

func queueInit() *queue {
	return &queue{}
}

func (q *queue) enQueue(num int) {
	q.S = append(q.S, num)
}

func (q *queue) push() int {
	if len(q.S) == 0 {
		return -1
	}
	num := q.S[0]
	q.S = q.S[1:]
	return num
}

func TestQueue(t *testing.T) {
	q := queueInit()
	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)
	q.enQueue(4)
	q.enQueue(5)

	fmt.Println(q.push())
	fmt.Println(q.push())
	fmt.Println(q.push())
	fmt.Println(q.push())
	fmt.Println(q.push())
	fmt.Println(q.push())
}

func enQueue(s []int, num int) []int {
	s = append(s, num)
	return s
}

func push(s []int) int {
	if len(s) == 0 {
		return -1
	}

	s = s[1:]
	return s[0]
}

func TestPush(t *testing.T) {
	s := []int{1, 4, 5, 3, 3}
	push(s)
	t.Log(s)
}

func radixSort(s []int) {
	// 个位数桶
	ma := make(map[int]*queue, 10)

	// 先把切片装入个位数桶
	for i, v := range s {
		if _, ok := ma[v%10]; ok {
			ma[v%10].enQueue(s[i])
		} else {
			ma[v%10] = queueInit()
			ma[v%10].enQueue(s[i])
		}

	}

	// 把个位数从0~9依次倒出来,组成一个新数组
	var sa []int
	for i := 0; i <= 9; i++ {
		if _, ok := ma[i]; ok {
			for {
				num := ma[i].push()
				if num == -1 {
					break
				}
				sa = append(sa, num)
			}
		}
	}

	// 把排好个位数的切片 再按照十位数装入桶内
	// 十位数桶
	mb := make(map[int]*queue, 10)
	for i, v := range s {
		if _, ok := mb[v/10%10]; ok {
			mb[v/10%10].enQueue(s[i])
		} else {
			mb[v/10%10] = queueInit()
			mb[v/10%10].enQueue(s[i])
		}
	}

	// 把十位数从0~9依次倒出来,组成一个新数组
	var sb []int
	for i := 0; i <= 9; i++ {
		if _, ok := mb[i]; ok {
			for {
				num := mb[i].push()
				if num == -1 {
					break
				}
				sb = append(sb, num)
			}
		}
	}

	fmt.Println(sb)
}

func TestRadixSort(t *testing.T) {
	s := []int{43, 24, 63, 73, 83, 13, 14, 74}
	radixSort(s)
}
