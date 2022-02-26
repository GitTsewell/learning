package linked_list

import (
	"fmt"
	"testing"
)

// 实现一个单向链表
type Node struct {
	Value interface{}
	Next  *Node
}

func signalListInit(v interface{}) *Node {
	return &Node{Value: v}
}

func (slf *Node) add(v interface{}) {
	if slf.Next != nil {
		slf.Next.add(v)
	} else {
		slf.Next = &Node{
			Value: v,
			Next:  nil,
		}
	}
}
func TestSignalList(t *testing.T) {
	list := signalListInit(1)
	list.add(2)
	list.add(3)
	list.add(4)

	fmt.Println(list)
}

// 实现一个双向链表
type TwoWayNode struct {
	Value interface{}
	Prev  *TwoWayNode
	Next  *TwoWayNode
}

func TwoWayNodeInit(v interface{}) *TwoWayNode {
	return &TwoWayNode{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}
}

func (slf *TwoWayNode) add(v interface{}) {
	if slf.Next != nil {
		slf.Next.add(v)
	} else {
		slf.Next = &TwoWayNode{
			Value: v,
			Prev:  slf,
			Next:  nil,
		}
	}
}

func TestTwoWayList(t *testing.T) {
	list := TwoWayNodeInit(1)
	list.add(2)
	list.add(3)
	list.add(4)

	fmt.Println(list)
}

// 单向链表反转
// A->B->C->D
// 设 n = nil
// 第一轮 n=&node{v : A.v , next : n}  反转之后第一个变成了最后一个 所以next 肯定是nil,所以一开始设n=nil
// 第二轮 n=&node{v : B.v , next : n} 把第二个B.next 指向刚才生成的A 实现反转
// 第三轮 n=&node{v : C.v , next : n}
func (slf *Node) nodeReverse() *Node {
	var n *Node
	cur := slf
	for cur != nil {
		n = &Node{Value: cur.Value, Next: n}
		cur = cur.Next
	}
	return n
}

func TestNodeReverse(t *testing.T) {
	list := signalListInit(1)
	list.add(2)
	list.add(3)
	list.add(4)

	r := list.nodeReverse()
	fmt.Println(r)
}

// 双向链表反转
// A<->B<->C<->D
// 设 n = nil
// 第一轮 n=&node{v : A.v , next : n} if n.next != nil {n.next.prev = n}  反转之后第一个变成了最后一个 所以next 肯定是nil,所以一开始设n=nil, prev就是下一轮生成的n ,这轮拿不到 下一轮再拿
// 第一轮 n=&node{v : B.v , next : n} if n.next != nil {n.next.prev = n} 把第二个B.next 指向刚才生成的A 然后让A.prev 指向B 也就是n.next.prev = n
// 第一轮 n=&node{v : B.v , next : n} if n.next != nil {n.next.prev = n}
func (slf *TwoWayNode) nodeReverse() *TwoWayNode {
	var n *TwoWayNode
	cur := slf
	for cur != nil {
		n = &TwoWayNode{Value: cur.Value, Next: n}
		if n.Next != nil {
			n.Next.Prev = n
		}

		cur = cur.Next
	}
	return n
}

func TestTwoWayNodeReverse(t *testing.T) {
	list := TwoWayNodeInit(1)
	list.add(2)
	list.add(3)
	list.add(4)

	r := list.nodeReverse()
	fmt.Println(r)
}
