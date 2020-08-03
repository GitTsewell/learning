package gotest

import "testing"

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(20) // 运行 Fib 函数 N 次
	}
}
