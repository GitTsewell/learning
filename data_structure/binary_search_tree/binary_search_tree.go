package binary_search_tree

import "fmt"

type Node struct {
	left  *Node
	right *Node
	val   int
}

func NewNode(val int) *Node {
	return &Node{val: val}
}

func (n *Node) Add(val int) *Node {
	if n == nil {
		n = NewNode(val)
		return n
	}

	if val < n.val {
		n.left = n.left.Add(val)
	} else {
		n.right = n.right.Add(val)
	}

	return n
}

func (n *Node) Contains(val int) bool {
	if n == nil {
		return false
	}

	if n.val == val {
		return true
	}

	if val < n.val {
		return n.left.Contains(val)
	} else {
		return n.right.Contains(val)
	}
}

//func (n *Node) Remove(val int) *Node {
//	if n == nil {
//		return n
//	}
//
//	// 先寻找
//	// 没找到 不做任何操作直接返回
//	// 找到后 判断有没有左子节点  有让左子节点中value最大的那个节点替代自己  如果没有 再判断有没有右子节点  如果有再让右子节点替代自己  都没有的话直接删除
//	if val < n.val {
//		return n.left.Remove(val)
//	} else if val > n.val {
//		return n.right.Remove(val)
//	} else if n.left != nil && n.right != nil {
//		n.left = n.left.left
//	} else if n.left != nil {
//		n = n.left
//	} else {
//		n = n.right
//	}
//
//	return n
//}

func (n *Node) Remove(val int) *Node {
	if n == nil {
		return n
	}

	if val < n.val {
		n.left = n.left.Remove(val)
	} else if val > n.val {
		n.right = n.right.Remove(val)
	} else if n.left != nil && n.right != nil {
		n.val = n.right.FindMin()
		n.right = n.right.Remove(n.val)
	} else if n.left != nil {
		n = n.left
	} else {
		n = n.right
	}
	return n
}

func (n *Node) FindMin() int {
	if n == nil {
		fmt.Println("tree is empty")
		return -1
	}
	if n.left == nil {
		return n.val
	} else {
		return n.left.FindMin()
	}
	//也可以直接用下面的方法
	//return t.FindMinNode().value
}
