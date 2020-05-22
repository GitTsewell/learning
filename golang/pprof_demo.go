package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

// 运行一段时间：fatal error: runtime: out of memory
//  go tool pprof http://localhost:6060/debug/pprof/heap
//top：显示正运行到某个函数goroutine的数量
//traces：显示所有goroutine的调用栈
//list：列出代码详细的信息。
func main() {
	// 开启pprof
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	tick := time.Tick(time.Second / 500)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
}
