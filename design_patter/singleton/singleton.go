package singleton

import "sync"

type S struct {
}

var singleton *S
var once sync.Once

func GetInstance() *S {
	once.Do(func() {
		singleton = &S{}
	})

	return singleton
}
