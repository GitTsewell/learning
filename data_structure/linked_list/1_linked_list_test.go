package linked_list

import (
	"fmt"
	"testing"
)

func TestSignalList(t *testing.T) {
	list := &Node{Value: 1}
	list.add(2)
	list.add(3)
	list.add(4)

	fmt.Println(list)
}

func TestTwoWayList(t *testing.T) {
	list := TwoWayNode{Value: 1}
	list.add(2)
	list.add(3)
	list.add(4)

	fmt.Println(list)
}
