package binary_tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	n := &Node{Value: 1}
	LeftInsert(n, 2)
	RightInsert(n, 3)
	LeftInsert(n.Left, 4)
	RightInsert(n.Left, 5)
	LeftInsert(n.Right, 6)
	RightInsert(n.Right, 7)

	fmt.Println(n)
}
