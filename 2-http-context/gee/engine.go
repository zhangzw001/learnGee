package gee

import "net/http"

type HandlerFunc func( *Context)

// 这里包装engine
// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
}
func NewEngine() *Engine {
	return &Engine{newRouter()}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// server 里面做什么呢?
	// 首先创建context
	c := newContext(w, req )
	// 然后 根据路由的规则调用对应的方法, 实现不同路由的处理
	e.handle(c)
}

func (e *Engine) Run(addr string) error  {
	return http.ListenAndServe(addr,e)
}
