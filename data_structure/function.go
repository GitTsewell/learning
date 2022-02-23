package data_structure

import (
	"math/rand"
	"time"
)

func RandSlice(n int) (s []int) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < n; i++ {
		s = append(s, rand.Intn(n))
	}

	return
}
