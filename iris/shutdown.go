package main

import (
	stdCtx "context"
	"sync"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	// 优雅的关闭程序
	serverWG := new(sync.WaitGroup)
	defer serverWG.Wait()

	iris.RegisterOnInterrupt(func() {
		serverWG.Add(1)
		defer serverWG.Done()

		ctx, cancel := stdCtx.WithTimeout(stdCtx.Background(), 20*time.Second)
		defer cancel()

		// 关闭所有主机
		app.Shutdown(ctx)
	})

	app.Get("/", func(ctx context.Context) {
		time.Sleep(5 * time.Second)
		ctx.JSON(iris.Map{"code": 1000, "data": "Welcome"})
	})

	app.Get("/hello", func(ctx context.Context) {
		ctx.JSON(iris.Map{"code": 1000, "data": "Hello Iris"})
	})

	app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler)
}
