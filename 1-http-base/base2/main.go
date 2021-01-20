package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type geeHandle struct{}

func (g geeHandle) ServeHTTP(response http.ResponseWriter, req  *http.Request) {
	matchedHeader,_ := regexp.MatchString(`^/header/.*`,req.URL.Path)
	switch  {
	case matchedHeader:
		fmt.Fprintf(response,"url.Path : %q\n",req.URL.Path)
		for h, v := range req.Header {
			fmt.Fprintf(response,"header[%q] : %q\n",h,v)
		}
	default :
		fmt.Fprintf(response,"url.Path : %q\n",req.URL.Path)
	}
}


func main() {
	// 如果一个结构体 geeHandle 实现了 Handler接口, 那么他就可以传参给ListenAndServe
	// 并且所有的请求全部会到执行 geeHandle.ServeHTTP
	// http.ListenAndServe
	// -> server.ListenAndServe
	// -> srv.Serve
	// -> go c.serve(connCtx)
	// -> serverHandler{c.server}.ServeHTTP(w, w.req)
	gee := &geeHandle{}
	http.ListenAndServe("localhost:8000",gee)

}
