package main

import (
	"log"
	"net/http"
	"org/jingtao8a/gee/context"
	"org/jingtao8a/gee/engine"
)

func indexRoute(ctx *context.Context) {
	//ctx.String(http.StatusOK, "%s", "yuxintao")
	ctx.JSON(http.StatusOK, map[string]int{"age": 24, "weight": 78})
}

func main() {
	engine := engine.NewEngine()
	engine.GET("/", indexRoute)
	log.Fatal(engine.Run(":9999"))
}
