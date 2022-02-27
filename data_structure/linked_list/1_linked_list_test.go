package linked_list

import (
	"fmt"
	"testing"
)

func TestSignalList(t *testing.T) {
	list := signalListInit()
	list.add(1)
	list.add(2)
	list.add(3)

	fmt.Println(list)
}

func TestTwoWayList(t *testing.T) {
	list := TwoWayNodeInit()
	list.add(2)
	list.add(3)
	list.add(4)

	fmt.Println(list)
}
