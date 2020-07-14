package hash

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	myHash := Constructor()
	myHash.Add(1)
	myHash.Add(2)
	fmt.Println(myHash.Contains(1))
	fmt.Println(myHash.Contains(3))
	myHash.Add(2)
	fmt.Println(myHash.Contains(2))
	myHash.Remove(2)
	fmt.Println(myHash.Contains(2))
}
