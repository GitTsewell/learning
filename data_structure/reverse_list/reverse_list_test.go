package reverse_list

import (
	"fmt"
	"testing"
)

func MakeList(val int) *ListNode {
	list := NewNode(0)
	node := list
	for i := 1; i <= val; i++ {
		node = node.Add(i)
	}
	return list
}

func TestReverseList(t *testing.T) {
	list := MakeList(10)
	reverse := ReverseList(list)
	fmt.Println(reverse)
}
