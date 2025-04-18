package engine

import (
	"log"
	"net/http"
	"strings"

	"org/jingtao8a/gee/router"
)

type RouterGroup struct {
	prefix      string
	middlewares []router.HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

type Engine struct {
	*RouterGroup
	router *router.Router
	groups []*RouterGroup
}

func NewEngine() *Engine {
	engine := &Engine{router: router.NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) Use(middlewares ...router.HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup) addRoute(method string, pattern string, handler router.HandlerFunc) {
	pattern = group.prefix + pattern
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.AddRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler router.HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler router.HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []router.HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	ctx := router.NewContext(w, r)
	ctx.Handlers = middlewares
	engine.router.Handle(ctx)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
