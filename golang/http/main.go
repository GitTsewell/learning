package main

func main() {
	//NewNetHttpServer()
	//NewFastHttpServer()
	//NewGinHttpServer()
	NewIrisHttpServer()
}

// net 197600             59351 ns/op              16 B/op          2 allocs/op
// fast 257317             48929 ns/op              16 B/op          2 allocs/op
// gin 140629             81024 ns/op              16 B/op          2 allocs/op
// iris 188812             64537 ns/op              16 B/op          2 allocs/op
