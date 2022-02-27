package linked_list

import (
	"fmt"
	"testing"
)

func TestNodeReverse(t *testing.T) {
	list := signalListInit()
	list.add(1)
	list.add(2)
	list.add(3)

	r := nodeReverse(list)
	fmt.Println(r)
}

func TestTwoWayNodeReverse(t *testing.T) {
	list := TwoWayNodeInit()
	list.add(2)
	list.add(3)
	list.add(4)

	r := TwoWayNodeReverse(list)
	fmt.Println(r)
}

func TestLinkedListPub(t *testing.T) {
	list1 := signalListInit()
	list1.add(1)
	list1.add(3)
	list1.add(5)
	list1.add(7)
	list1.add(9)

	list2 := signalListInit()
	list2.add(2)
	list2.add(3)
	list2.add(4)
	list2.add(5)
	list2.add(6)
	list2.add(7)

	pub := LinkedListPub(list1, list2)
	fmt.Println(pub)
}
