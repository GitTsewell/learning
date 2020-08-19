package LRU

import "testing"

func TestLruCache(t *testing.T) {
	lru := Constructor(10)

	for i := 1; i <= 10; i++ {
		lru.Put(i, i)
	}
	t.Log(lru.Get(5))
}
