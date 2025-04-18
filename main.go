package main

import (
	"log"
	"net/http"
	"org/jingtao8a/gee/engine"
	"org/jingtao8a/gee/logger"
	"org/jingtao8a/gee/router"
)

func main() {
	engine := engine.NewEngine()
	engine.Use(logger.Logger())
	engine.GET("/", func(ctx *router.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	v1 := engine.Group("/v1")
	{
		v1.GET("/test", func(ctx *router.Context) {
			ctx.String(http.StatusOK, "this is test")
		})
		v1.GET("/hello/:name", func(ctx *router.Context) {
			ctx.String(http.StatusOK, "hello world %s", ctx.Params["name"])
		})
	}
	v2 := engine.Group("/v2")
	{
		v2.GET("/assets/:name/*filepath", func(ctx *router.Context) {
			ctx.JSON(http.StatusOK, ctx.Params)
		})
		v2.GET("/assets/*filepath", func(ctx *router.Context) {
			ctx.JSON(http.StatusOK, ctx.Params)
		})
	}
	log.Fatal(engine.Run(":9999"))
}
