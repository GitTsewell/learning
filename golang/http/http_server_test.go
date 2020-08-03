package main

import (
	"net/http"
	"testing"
)

func BenchmarkNewHttpServer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i <= b.N; i++ {
		Request(http.MethodGet, "http://127.0.0.1:30001/ping", nil)
	}
}

//go test -bench=BenchmarkNewNetHttpServer -benchmem -benchtime=10s
//go test -v  demo_test.go -test.bench Demo3 -test.run Demo3
