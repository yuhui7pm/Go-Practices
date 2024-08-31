package main

import (
	"github.com/kataras/iris/v12"
	"log"
	"path/filepath"
	"runtime"
)

func getCurrentDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Could not get current file path")
	}
	return filepath.Dir(filename)
}
func main() {
	app := iris.New()

	htmlEngine := iris.HTML(getCurrentDir(), ".html")

	app.RegisterView(htmlEngine)

	// 定义一个 GET 请求的路由
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello from Iris!")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "this is title")
		ctx.ViewData("Content", "this is content")
		err := ctx.View("hello.html")
		if err != nil {
			log.Println("Error rendering view:", err)
			return
		}

		log.Println("After rendering view...")
	})
	// 启动服务器
	app.Run(iris.Addr(":8080"))
}
