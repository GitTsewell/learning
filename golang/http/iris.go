package main

import (
	"github.com/kataras/iris"
)

func NewIrisHttpServer() {
	iris.New()
	app := iris.New()

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Run(iris.Addr(":30001"))
}
