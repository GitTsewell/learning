package linked_list

// 单向链表反转
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

// LinkedListPub 打印两个有序链表的公共部分
// 比如 1->3->5->7->9   和 2->3->4->5->6->7  打印出3,5,7
// 思路: 两个链表遍历 因为是有序的,所以谁小谁移动,另一个不动,如果相等,就追加进数组,然后两个都移动
func LinkedListPub(list1, list2 *Node) (data []interface{}) {
	if list1 == nil || list2 == nil {
		return
	}

	for list1.Next != nil && list2.Next != nil {
		if list1.Next.Value.(int) == list2.Next.Value.(int) {
			data = append(data, list1.Next.Value)
			list1.Next = list1.Next.Next
			list2.Next = list2.Next.Next
		} else if list1.Next.Value.(int) < list2.Next.Value.(int) {
			list1.Next = list1.Next.Next
		} else {
			list2.Next = list2.Next.Next
		}
	}

	return
}

// 判断一个链表是否会回文结构,回文结构 1->2->3->2->1  or  1->20->8->8->20->1 要求时间复杂度O(n)
// 思路:1 链表反转 然后比较遍历 比较是否一致
// 2. 使用栈结构,依次入栈,然后遍历和出栈结构比较,和链表反转思路类型,不过这里把这个反转链表的结果变成了一个栈结构,利用栈后进先出的特点,让首位比较
// 3. 优化栈结构,只使用后半部分入栈,然后和前半部分去比较,可以节约N/2 空间,然后这个方法的问题就是如果找到中心对称点的问题,可以使用快慢指针,当快指针
// 走完的时候,慢指针就是指向的中心对称点,还要判断奇数偶数的问题,如果是基数是刚好有一个中心对称点的,如果是偶数,慢指针停在中心对称右边或右边一个数,可以自己控制

func ListPalindromeV1(head *Node) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return true
	}

	node := *head
	reverseNode := nodeReverse(&node)

	for head.Next != nil && reverseNode.Next != nil {
		if head.Value != reverseNode.Value {
			return false
		}

		head = head.Next
		reverseNode = reverseNode.Next
	}

	return true
}
