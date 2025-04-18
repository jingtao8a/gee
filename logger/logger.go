package logger

import (
	"log"
	"org/jingtao8a/gee/router"
	"time"
)

func Logger() router.HandlerFunc {
	return func(ctx *router.Context) {
		t := time.Now()
		ctx.Next()
		log.Printf("[%d] %s in %v", ctx.StatusCode, ctx.Method+" "+ctx.Req.RequestURI, time.Since(t))
	}
}
