package main

import (
	"sync"
	"testing"
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
