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
	v1 := engine.Group("/v1")
	{
		v1.GET("/test", func(ctx *context.Context) {
			ctx.String(http.StatusOK, "this is test")
		})
		v1.GET("/hello/:name", func(ctx *context.Context) {
			ctx.String(http.StatusOK, "hello world %s", ctx.Params["name"])
		})
	}
	v2 := engine.Group("/v2")
	{
		v2.GET("/assets/*filepath", func(ctx *context.Context) {
			ctx.JSON(http.StatusOK, ctx.Params["filepath"])
		})
		v2.GET("/assets/:name/*filepath", func(ctx *context.Context) {
			ctx.JSON(http.StatusOK, ctx.Params)
		})
	}
	log.Fatal(engine.Run(":9999"))
}
