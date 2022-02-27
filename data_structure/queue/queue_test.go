package queue

import (
	"fmt"
	"testing"
)

func TestArrQueue(t *testing.T) {
	q := ArrQueueInit()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}

func TestListQueue(t *testing.T) {
	q := ListQueueInit()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
