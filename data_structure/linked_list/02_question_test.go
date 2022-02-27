package linked_list

import (
	"fmt"
	"testing"
)

func TestNodeReverse(t *testing.T) {
	list := &Node{Value: 1}
	list.add(2)
	list.add(3)

	r := nodeReverse(list)
	fmt.Println(r)
}

func TestTwoWayNodeReverse(t *testing.T) {
	list := &TwoWayNode{Value: 1}
	list.add(2)
	list.add(3)
	list.add(4)

	r := TwoWayNodeReverse(list)
	fmt.Println(r)
}

func TestLinkedListPub(t *testing.T) {
	list1 := &Node{Value: 1}
	list1.add(3)
	list1.add(5)
	list1.add(7)
	list1.add(9)

	list2 := &Node{Value: 2}
	list2.add(3)
	list2.add(4)
	list2.add(5)
	list2.add(6)
	list2.add(7)

	pub := LinkedListPub(list1, list2)
	fmt.Println(pub)
}

func TestListPalindromeV1(t *testing.T) {
	list1 := &Node{Value: 1}
	list1.add(3)
	list1.add(5)
	list1.add(7)
	list1.add(9)

	fmt.Println(ListPalindromeV1(list1))

	list2 := &Node{Value: 1}
	list2.add(3)
	list2.add(5)
	list2.add(3)
	list2.add(1)
	fmt.Println(ListPalindromeV1(list2))

	list3 := &Node{Value: 20}
	list3.add(4)
	list3.add(4)
	list3.add(20)
	fmt.Println(ListPalindromeV1(list3))
}
