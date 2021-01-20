package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type H map[string]interface{}

// 这里包装Context
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req *http.Request
	Header http.Header

	// request info
	Path string
	Method string
	Params map[string]string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		Header: req.Header,
	}
}

func (c *Context) Param(key string) string {
	value , _ := c.Params[key]
	return value
}

func (c *Context) PostForm(k string) string{
	return c.Req.FormValue(k)
}

func (c *Context) Query(k string)  string {
	return c.Req.URL.Query().Get(k)
}

func (c *Context) Status(code int ) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(k string , v string ) {
	c.Writer.Header().Set(k,v )
}

func (c *Context) String(code int , format string, value ...interface{} ) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format,value...)))
}

func (c *Context) JSON(code int , obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
