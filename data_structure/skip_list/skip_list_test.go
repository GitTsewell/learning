package skip_list

import (
	"sync"
	"testing"
)

func TestSkipList_Add(t *testing.T) {
	skip := Constructor()
	for i := 0; i < 10000; i++ {
		skip.Add(i + 1)
	}

	t.Log(skip.Search(6200))
}

func TestSkipList_Add_total(t *testing.T) {
	var total int
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i <= 1000; i++ {
		go func() {
			skip := Constructor()
			for i := 0; i < 10000; i++ {
				skip.Add(i + 1)
			}
			time, _ := skip.Search(6200)
			total += time
			wg.Done()
		}()
	}

	wg.Wait()
	t.Log(total / 1000)
}
