package stack

// ArrStack 使用数组来实现栈结构 init() push() pop()
type ArrStack struct {
	Head []interface{}
}

func ArrStackInit() *ArrStack {
	return &ArrStack{Head: []interface{}{}}
}

func (a *ArrStack) Push(v interface{}) {
	a.Head = append(a.Head, v)
}

func (a *ArrStack) Pop() interface{} {
	if len(a.Head) == 0 {
		return nil
	}

	data := a.Head[len(a.Head)-1]
	a.Head = a.Head[:len(a.Head)-1]
	return data
}

// Stack 使用链表实现stack 结构  和用list 实现queue类似,nil->A->B->C->nil ,不过push的时候不是追加的链尾,而是插入到nil和A之间
type Stack struct {
	Head *Node
}

type Node struct {
	Value interface{}
	Next  *Node
}

func StackInit() *Stack {
	return &Stack{
		Head: &Node{},
	}
}

func (s *Stack) Push(v interface{}) {
	if s == nil {
		return
	}

	node := &Node{
		Value: v,
		Next:  nil,
	}

	if s.Head.Next == nil {
		s.Head.Next = node
	} else {
		tmp := s.Head.Next
		node.Next = tmp
		s.Head.Next = node

	}
}

func (s *Stack) Pop() interface{} {
	if s == nil || s.Head.Next == nil {
		return nil
	}

	data := s.Head.Next.Value
	s.Head.Next = s.Head.Next.Next
	return data
}
