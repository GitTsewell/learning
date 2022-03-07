package binary_tree

import "testing"

func MakeTree() *Node {
	n := &Node{Value: 1}
	LeftInsert(n, 2)
	RightInsert(n, 3)
	LeftInsert(n.Left, 4)
	RightInsert(n.Left, 5)
	LeftInsert(n.Right, 6)
	RightInsert(n.Right, 7)
	return n
}

func TestPreorderTraversal(t *testing.T) {
	n := MakeTree()
	PreorderTraversal(n)
}

func TestInorderTraversal(t *testing.T) {
	n := MakeTree()
	InorderTraversal(n)
}

func TestPostorderTraversal(t *testing.T) {
	n := MakeTree()
	PostorderTraversal(n)
}

func TestPreorderTraversalV2(t *testing.T) {
	n := MakeTree()
	PreorderTraversalV2(n)
}

func TestInorderTraversalV2(t *testing.T) {
	n := MakeTree()
	InorderTraversalV2(n)
}
