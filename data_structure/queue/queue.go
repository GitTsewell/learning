package queue

type ArrQueue struct {
	Queue []interface{}
}

// ArrQueueInit 使用数组实现一个队列 init() push()  pop()
func ArrQueueInit() *ArrQueue {
	return &ArrQueue{Queue: []interface{}{}}
}

func (a *ArrQueue) Push(v interface{}) {
	a.Queue = append(a.Queue, v)
}

func (a *ArrQueue) Pop() (v interface{}) {
	if len(a.Queue) == 0 {
		return -1
	}

	v = a.Queue[0]
	a.Queue = a.Queue[1:]
	return
}

// Queue 使用一个链表 实现queue
// nil->A->B->C->D->nil  一开始head和tail->nil 头的nil是不会变的,pop弹出A,push的时候把tail->A
type Queue struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value interface{}
	Next  *Node
}

func ListQueueInit() *Queue {
	node := &Node{}
	return &Queue{
		Head: node,
		Tail: node,
	}
}

func (q *Queue) Push(v interface{}) {
	if q == nil || q.Tail == nil {
		return
	}

	node := &Node{Value: v, Next: nil}
	q.Tail.Next = node
	q.Tail = q.Tail.Next
}

func (q *Queue) Pop() interface{} {
	if q == nil || q.Head == nil || q.Head == q.Tail {
		return nil
	}

	if q.Head.Next == q.Tail {
		q.Tail = q.Head
	}

	n := q.Head.Next.Value
	q.Head.Next = q.Head.Next.Next
	return n
}
