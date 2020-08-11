package reverse_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func (slf *ListNode) Add(val int) *ListNode {
	node := NewNode(val)
	slf.Next = node
	return node
}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for nil != cur {
		// 1.（保存一下前进方向）保存下一跳
		temp := cur.Next
		// 2.斩断过去,不忘前事
		cur.Next = pre
		// 3.前驱指针的使命在上面已经完成，这里需要更新前驱指针
		pre = cur
		// 当前指针的使命已经完成，需要继续前进了
		cur = temp
	}
	return pre
}
