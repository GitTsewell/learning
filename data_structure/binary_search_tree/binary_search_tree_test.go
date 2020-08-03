package binary_search_tree

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNode_Add(t *testing.T) {
	node := NewNode(0)

	node.Add(100)
	for i := 0; i <= 5; i++ {
		node.Add(rand.Intn(1000))
	}
	node.Remove(100)
	fmt.Println(node)
}
