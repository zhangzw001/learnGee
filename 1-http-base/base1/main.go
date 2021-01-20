package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/header", headerHandler)
	// ListenAndServe 会调用 ServeHTTP()
	http.ListenAndServe("localhost:8000", nil)
}

func indexHandler(response http.ResponseWriter, req *http.Request) {
	// 带引号输出
	fmt.Fprintf(response, "url.Path : %q\n", req.URL.Path)
	// 不带引号
	//fmt.Fprintf(response,"url.Path : %s\n",req.URL.Path)

}

func headerHandler(response http.ResponseWriter, req *http.Request) {
	// 输出这个  type Header map[string][]string
	for h, v := range req.Header {
		fmt.Fprintf(response, "header[%q] : %q\n", h, v)
	}

}
