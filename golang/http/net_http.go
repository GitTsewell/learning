package main

import "net/http"

func netHttpHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func NewNetHttpServer() {
	http.HandleFunc("/ping", netHttpHandle)
	if err := http.ListenAndServe(":30001", nil); err != nil {
		panic(err.Error())
	}
}
