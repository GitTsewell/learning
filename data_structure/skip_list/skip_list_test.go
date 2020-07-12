package skip_list

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSkiplist_Add(t *testing.T) {
	skip := Constructor()
	skip.Add(3901)
	for i := 0; i < 10; i++ {
		skip.Add(rand.Intn(20000))
	}
	fmt.Println(skip.Search(3901))
	fmt.Println(skip.Erase(3901))
	fmt.Println(skip)
}
