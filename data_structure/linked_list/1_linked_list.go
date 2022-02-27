package linked_list

// Node 实现一个单向链表
type Node struct {
	Value interface{}
	Next  *Node
}

func signalListInit(v interface{}) *Node {
	return &Node{Value: v}
}

func (n *Node) add(v interface{}) {
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

func TwoWayNodeInit(v interface{}) *TwoWayNode {
	// sds
	return &TwoWayNode{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}
}

func (t *TwoWayNode) add(v interface{}) {
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

// 单向链表反转
// A->B->C->D
// 设 n = nil
// 第一轮 n=&node{v : A.v , next : n}  反转之后第一个变成了最后一个 所以next 肯定是nil,所以一开始设n=nil
// 第二轮 n=&node{v : B.v , next : n} 把第二个B.next 指向刚才生成的A 实现反转
// 第三轮 n=&node{v : C.v , next : n}
func (n *Node) nodeReverse() *Node {
	var node *Node
	cur := n
	for cur != nil {
		node = &Node{Value: cur.Value, Next: node}
		cur = cur.Next
	}
	return node
}

// A->B->C->D
// 先用一个临时变量把(B->C->D)保存下来
// 把B->A 也就是newHead
// 把B赋值给newHead
// 把head 接回临时变量的(B->C->D)
func nodeReverse(head *Node) *Node {
	var newHead *Node

	for head != nil {
		node := head.Next
		head.Next = newHead
		newHead = head
		head = node
	}
	return newHead
}

// 双向链表反转
// A<->B<->C<->D
// 设 n = nil
// 第一轮 n=&node{v : A.v , next : n} if n.next != nil {n.next.prev = n}  反转之后第一个变成了最后一个 所以next 肯定是nil,所以一开始设n=nil, prev就是下一轮生成的n ,这轮拿不到 下一轮再拿
// 第一轮 n=&node{v : B.v , next : n} if n.next != nil {n.next.prev = n} 把第二个B.next 指向刚才生成的A 然后让A.prev 指向B 也就是n.next.prev = n
// 第一轮 n=&node{v : B.v , next : n} if n.next != nil {n.next.prev = n}
func (t *TwoWayNode) nodeReverse() *TwoWayNode {
	var n *TwoWayNode
	cur := t
	for cur != nil {
		n = &TwoWayNode{Value: cur.Value, Next: n}
		if n.Next != nil {
			n.Next.Prev = n
		}

		cur = cur.Next
	}
	return n
}

func TwoWayNodeReverse(head *TwoWayNode) *TwoWayNode {
	var newHead *TwoWayNode
	for head != nil {
		node := head.Next
		head.Next = newHead
		// head.Prev
		newHead = head
		head = node
		newHead.Prev = head
	}

	return newHead
}
