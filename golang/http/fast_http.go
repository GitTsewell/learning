package main

import (
	"errors"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
)

func fastHttpHandle(ctx *fasthttp.RequestCtx) {
	ctx.Write([]byte("pong"))
	//fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
}

func NewFastHttpServer() {
	fasthttp.ListenAndServe(":30001", fastHttpHandle)
}

func Request(method, url string, data []byte) (result []byte, err error) {
	req := fasthttp.AcquireRequest()
	rsp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(rsp)
		fasthttp.ReleaseRequest(req)
	}()

	req.SetRequestURI(url)
	if len(data) > 0 {
		req.SetBody(data)
	}

	req.Header.SetMethod(strings.ToUpper(method))

	//fmt.Printf("http request method : %s , url : %s , data : %s \n", method, url, data)
	if err := fasthttp.DoTimeout(req, rsp, time.Second*30); err != nil {
		return nil, errors.New("http read response err")
	}
	//fmt.Printf("http response : %s \n", rsp.Body())

	return rsp.Body(), nil
}
