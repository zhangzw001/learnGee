package main

import "C"
import (
	"fmt"
	"github.com/zhangzw001/learnGee/1-http-base/base3/gee"
	"net/http"
)


func main() {
	group := gee.New()
	group.GET("/header",headerHandler)
	group.POST("/add",addHandler)
	http.ListenAndServe("localhost:8000",group)
}

func headerHandler(w http.ResponseWriter, req *http.Request) {
	// 输出这个  type Header map[string][]string
	for h, v := range req.Header {
		fmt.Fprintf(w,"header[%q] : %q\n",h,v)
	}
}


func addHandler(w http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	fmt.Fprintln(w, req.PostForm)
}
