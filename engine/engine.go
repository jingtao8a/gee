package engine

import (
	"log"
	"net/http"

	"org/jingtao8a/gee/context"
	"org/jingtao8a/gee/router"
)

type Engine struct {
	router *router.Router
}

func NewEngine() *Engine {
	return &Engine{router: router.NewRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler router.HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.AddRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler router.HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler router.HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.NewContext(w, r)
	engine.router.Handle(ctx)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
