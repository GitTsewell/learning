package linked_list

// Node 实现一个单向链表
type Node struct {
	Value interface{}
	Next  *Node
}

func signalListInit() *Node {
	return &Node{}
}

func (n *Node) add(v interface{}) {
	if n == nil {
		return
	}

	if n.Next != nil {
		n.Next.add(v)
	} else {
		n.Next = &Node{
			Value: v,
			Next:  nil,
		}
	}
}

// TwoWayNode 实现一个双向链表
type TwoWayNode struct {
	Value interface{}
	Prev  *TwoWayNode
	Next  *TwoWayNode
}

func TwoWayNodeInit() *TwoWayNode {
	// sds
	return &TwoWayNode{}
}

func (t *TwoWayNode) add(v interface{}) {
	if t == nil {
		return
	}

	if t.Next != nil {
		t.Next.add(v)
	} else {
		t.Next = &TwoWayNode{
			Value: v,
			Prev:  t,
			Next:  nil,
		}
	}
}
