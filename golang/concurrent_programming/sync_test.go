package main

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestSyncMutex(t *testing.T) {
	mutex := sync.Mutex{}
	mutex.Lock()
	mutex.Unlock()
}

func TestSyncRWMutex(t *testing.T) {
	rw := sync.RWMutex{}
	rw.RLock()
	go rw.Lock()
	rw.RUnlock()
	rw.Unlock()
}

func worker(i int) {
	fmt.Println("worker: ", i)
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}
	wg.Wait()
}

type WaitGroup struct {
	noCopy1 noCopy1

	// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
	// 64-bit atomic operations require 64-bit alignment, but 32-bit
	// compilers do not ensure it. So we allocate 12 bytes and then use
	// the aligned 8 bytes in them as state, and the other 4 as storage
	// for the sema.
	state1 [3]uint32
}

type noCopy1 struct{}

func (*noCopy1) Lock()   {}
func (*noCopy1) Unlock() {}

// Lock is a no-op used by -copylocks checker from `go vet`.

func TestNoCopy(t *testing.T) {
	w := WaitGroup{}
	w1 := w
	fmt.Println(w1)
}

func TestSyncOnce(t *testing.T) {
	var (
		o  sync.Once
		wg sync.WaitGroup
	)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			o.Do(func() {
				fmt.Println("once", i)
			})
		}(i)
	}

	wg.Wait()
}

func TestErrGroup(t *testing.T) {
	var g errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				_ = resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	}
}
