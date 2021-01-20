package gee

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

// 这里包装engine
// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
	*router
	groups []*RouterGroup	// 用来存在所有的group
}

func NewEngine() *Engine {
	engine:= &Engine{router:newRouter()}
	engine.RouterGroup = &RouterGroup{engine:engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// server 里面做什么呢?
	// 处理中间件
	var handlers []HandlerFunc
	// Use的时候会添加中间件到groups的handlers 列表
	for _, group := range e.groups {
		//当我们接收到一个具体请求时，要判断该请求适用于哪些中间件，在这里我们简单通过 URL 的前缀来判断。得到中间件列表后，赋值给 c.handlers。
		// 只要是匹配到 group的前缀, 就需要添加这个前缀的group的所有中间件
		if strings.HasPrefix(req.URL.Path , group.prefix) {
			//log.Printf("Engine.ServeHTTP URL.Path : %s, group.prefix : %s\n",req.URL.Path , group.prefix)
			//log.Printf("Engine.ServeHTTP len(handlers) : %d\n",len(handlers))
			handlers = append(handlers, group.handlers...)
		}
	}
	// 首先创建context
	c := newContext(w, req )
	c.handlers = handlers
	// 然后 根据路由的规则调用对应的方法, 实现不同路由的处理
	e.router.handle(c)
}

func (e *Engine) Run(addr string) error  {
	return http.ListenAndServe(addr,e)
}


