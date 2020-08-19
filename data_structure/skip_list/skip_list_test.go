package skip_list

import (
	"testing"
)

func TestSkipList_Add(t *testing.T) {
	skip := Constructor()
	for i := 0; i < 10000; i++ {
		skip.Add(i + 1)
	}

	t.Log(skip.Search(6200))
}
