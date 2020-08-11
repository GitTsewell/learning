package skip_list

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSkiplist_Add(t *testing.T) {
	skip := Constructor()
	for i := 0; i < 10000000; i++ {
		skip.Add(rand.Intn(100000000))
	}
	fmt.Println(skip)
}
