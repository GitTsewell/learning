package linked_list

import (
	"fmt"
	"testing"
)

func TestSignalList(t *testing.T) {
	list := signalListInit(1)
	list.add(2)
	list.add(3)
	list.add(4)

	fmt.Println(list)
}

func TestTwoWayList(t *testing.T) {
	list := TwoWayNodeInit(1)
	list.add(2)
	list.add(3)
	list.add(4)

	fmt.Println(list)
}

func TestNodeReverse(t *testing.T) {
	list := signalListInit(1)
	list.add(2)
	list.add(3)
	list.add(4)

	r := nodeReverse(list)
	fmt.Println(r)
}

func TestTwoWayNodeReverse(t *testing.T) {
	list := TwoWayNodeInit(1)
	list.add(2)
	list.add(3)
	list.add(4)

	r := TwoWayNodeReverse(list)
	fmt.Println(r)
}
