package binary_tree

import (
	"fmt"
	"learning/data_structure/stack"
)

// 思路  二叉树如下
//					1
//			2				3
//		4		5		6		7
//	nil	nil	nil	nil	nil	nil	nil	nil
// 比如我们这样一个二叉树,按照PreorderTraversal()调用得到的栈链条是
// 12444555213666377731 先序遍历的顺序是 头,左,右  由调用顺序得知,可以在n第一次出现的时候打印n的值 1245367就是先序遍历 ,也就是n节点马上准备入栈之前
// 原理很简单,因为我们是二叉树结构,有左右两个节点,如果我们采用先寻找完左节点,然后回到头节点,再寻找右节点的顺序,那每个不为nil的节点肯定会入栈和出栈两次
// 比如我们遍历2节点 顺序是 2入栈 4入栈 nil 4出栈 nil 4入栈 nil 4出栈 2出栈 2入栈 5入栈 nil 5出栈 5入栈 nil 5出栈 2出栈
// 所以我们在node入栈之前 打印n的值,天然就是一个先序遍历

// 中序遍历 顺序是 左,头,右 ,由调用顺序得知,可以在n第二次出现的时候打印n的值就是中序遍历 4521637就是中序遍历
// 也就是node节点第一次出栈完成,准备第二次入栈时,打印n的值

// 后续遍历  左,右,头,在n第三次出现的时候打印n的值就是后续遍历
// 也就是node节点第二次出栈完成的时候

// PreorderTraversal 二叉树的先序遍历 递归版本 顺序是 头,左,右
func PreorderTraversal(n *Node) {
	if n == nil {
		return
	}

	fmt.Println(n.Value)

	// 调用node.left 递归,node节点入栈
	PreorderTraversal(n.Left)

	// node.left递归完成,node出栈
	// 调用node.right递归,node又入栈
	PreorderTraversal(n.Right)
	// 调用node.right递归完成,node出栈
}

// InorderTraversal 二叉树的中序遍历 递归版本 顺序是 左 头 右
func InorderTraversal(n *Node) {
	if n == nil {
		return
	}
	InorderTraversal(n.Left)
	fmt.Println(n.Value)
	InorderTraversal(n.Right)
}

// PostorderTraversal 二叉树的后续遍历 递归版本  顺序是 左 右 头
func PostorderTraversal(n *Node) {
	if n == nil {
		return
	}

	PostorderTraversal(n.Left)
	PostorderTraversal(n.Right)
	fmt.Println(n.Value)
}

// PreorderTraversalV2 二叉树的先序遍历 非递归版本 顺序是 头,左,右
// 思路 手动实现一个栈,用来替代系统栈
func PreorderTraversalV2(n *Node) {
	stacks := stack.StackInit()
	stacks.Push(n)

	for {
		node := stacks.Pop()
		if node == nil {
			break
		}
		fmt.Println(node.(*Node).Value)

		// 顺序是头 左 右,所以右边先入栈
		if node.(*Node).Right != nil {
			stacks.Push(node.(*Node).Right)
		}

		if node.(*Node).Left != nil {
			stacks.Push(node.(*Node).Left)
		}
	}
}

// InorderTraversalV2 中序遍历  左 头 右
// 在第一次出栈的时候打印
func InorderTraversalV2(n *Node) {
	tmpStacks := stack.StackInit()
	tmpStacks.Push(n)
	stacks := stack.StackInit()

	for {
		node := tmpStacks.Pop()
		if node == nil {
			break
		}
		stacks.Push(node)
		if node.(*Node).Right != nil {
			tmpStacks.Push(node.(*Node).Right)
		}

		if node.(*Node).Left != nil {
			tmpStacks.Push(node.(*Node).Left)
		}
	}

	for {
		node := stacks.Pop()
		if node == nil {
			break
		}

		fmt.Println(node.(*Node).Value)
	}
}
