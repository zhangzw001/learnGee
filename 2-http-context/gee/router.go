package gee

import (
	"net/http"
)

// 这里包装router
type RouterGroup struct {
	handlers map[string]HandlerFunc
}

func newRouter() *RouterGroup {
	return &RouterGroup{handlers: make(map[string]HandlerFunc)}
}


func (r *RouterGroup) handle(c *Context) {
	// 路由里面就是根据http请求的url 找到对应的handlerFunc 然后调用即可
	// 1. 拼接router的key
	key := c.Method+"-"+ c.Path
	// 2. 从map里面取到handler
	handler , ok := e.handlers[key]
	if ok {
		// 3. 调用方法
		handler(c)
	}else {
		c.String(http.StatusNotFound,"404 not found : %s\n",c.Path)
	}
}

//
func (r *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc ) {
	//log.Printf("Route %4s - %s", method, pattern)
	e.handlers[method+"-"+pattern] = handler
}

// 这里GET方法写在 router 里面
// engine 是包含 router
func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodGet,pattern,handler)
}

func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	e.addRoute(http.MethodPost,pattern,handler)
}
