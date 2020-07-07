package skip_list

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewNode(t *testing.T) {
	list := NewSkipList()
	for i := 0; i < 20; i++ {
		list.Insert(rand.Intn(100))
	}
	list.Print()

	fmt.Println("\n--------------------------------------")

	list.Delete(10)
	list.Print()

	fmt.Println("\n--------------------------------------")
}
