package context

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// 查询请求表单中key对应的value
func (ctx *Context) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

// 查找URL参数中key对应的value
func (ctx *Context) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

// 往Reponse填写状态码
func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	encoder := json.NewEncoder(ctx.Writer)
	if err := encoder.Encode(obj); err != nil {
		log.Printf("json encode error: %v", err)
	}
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.SetHeader("Content-Type", "application/octet-stream")
	ctx.Status(code)
	ctx.Writer.Write(data)
}

func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	ctx.Writer.Write([]byte(html))
}
