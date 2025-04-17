package router

import (
	"net/http"
	"org/jingtao8a/gee/context"
)

type HandlerFunc func(ctx *context.Context)

type Router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc)}
}

func (r *Router) AddRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *Router) Handle(ctx *context.Context) {
	key := ctx.Method + "-" + ctx.Path
	if handler, ok := r.handlers[key]; ok {
		handler(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 not found %s", ctx.Path)
	}
}
