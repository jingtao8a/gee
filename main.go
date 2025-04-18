package main

import (
	"log"
	"net/http"
	"org/jingtao8a/gee/context"
	"org/jingtao8a/gee/engine"
)

func main() {
	engine := engine.NewEngine()
	engine.GET("/", func(ctx *context.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	engine.GET("/test", func(ctx *context.Context) {
		ctx.String(http.StatusOK, "this is test")
	})
	engine.GET("/hello/:name", func(ctx *context.Context) {
		ctx.String(http.StatusOK, "hello world %s", ctx.Params["name"])
	})
	engine.GET("/assets/*filepath", func(ctx *context.Context) {
		ctx.JSON(http.StatusOK, ctx.Params["filepath"])
	})
	engine.GET("/assets/:name/*filepath", func(ctx *context.Context) {
		ctx.JSON(http.StatusOK, ctx.Params)
	})
	log.Fatal(engine.Run(":9999"))
}
