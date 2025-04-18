package router

import (
	"net/http"
	"org/jingtao8a/gee/tire"
	"org/jingtao8a/gee/util"
	"strings"
)

type HandlerFunc func(ctx *Context)

type Router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*tire.Node // 根据 GET  POST PUT  DELETE 分为不同的前缀树
}

func NewRouter() *Router {
	return &Router{handlers: make(map[string]HandlerFunc), roots: make(map[string]*tire.Node)}
}

func (r *Router) AddRoute(method string, pattern string, handler HandlerFunc) {
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &tire.Node{}
	}

	parts := util.ParsePattern(pattern)
	pattern = "/" + strings.Join(parts, "/")
	key := method + "-" + pattern

	r.roots[method].Insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) GetRoute(method string, pattern string) (*tire.Node, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	searchParts := util.ParsePattern(pattern)
	n := root.Search(searchParts, 0)
	if n == nil {
		return nil, nil
	}

	parts := util.ParsePattern(n.Pattern)

	params := make(map[string]string)
	for i := 0; i < len(parts); i++ {
		part := parts[i]
		if part[0] == ':' {
			params[part[1:]] = searchParts[i]
		} else if part[0] == '*' {
			params[part[1:]] = strings.Join(searchParts[i:], "/")
		}
	}
	return n, params
}

func (r *Router) Handle(ctx *Context) {
	n, params := r.GetRoute(ctx.Method, ctx.Path)
	if n != nil {
		ctx.Params = params
		key := ctx.Method + "-" + n.Pattern
		ctx.Handlers = append(ctx.Handlers, r.handlers[key])
	} else {
		ctx.Handlers = append(ctx.Handlers, func(ctx *Context) { ctx.String(http.StatusNotFound, "404 not found %s", ctx.Path) })
	}
	ctx.Next()
}
