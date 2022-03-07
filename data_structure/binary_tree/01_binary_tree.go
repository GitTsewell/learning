package binary_tree

type Node struct {
	Value interface{}
	Left  *Node
	Right *Node
}

func LeftInsert(n *Node, v interface{}) {
	n.Left = &Node{Value: v}
}

func RightInsert(n *Node, v interface{}) {
	n.Right = &Node{Value: v}
}
