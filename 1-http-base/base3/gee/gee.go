package gee

import (
	"fmt"
	"net/http"
)

//https://geektutu.com/post/gee-day1.html
type HandlerFunc func(http.ResponseWriter, *http.Request)

type RouterGroup struct {
	router map[string]HandlerFunc
}

func (r *RouterGroup) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 路由里面就是根据http请求的url 找到对应的handlerFunc 然后调用即可
	// 1. 拼接router的key
	key := req.Method+"-"+ req.URL.Path
	// 2. 从map里面取到handler
	handler , ok := e.router[key]
	if ok {
		// 3. 调用方法
		handler(w, req)
	}else {
		fmt.Fprintf(w,"404 not found %s\n",req.URL)
	}
}

func New() *RouterGroup {
	return &RouterGroup{router: make(map[string]HandlerFunc)}
}

func (r *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc ) {
	e.router[method+"-"+pattern] = handler
}


func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET",pattern,handler)
}

func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST",pattern,handler)
}


