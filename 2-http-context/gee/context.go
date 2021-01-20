package gee

import (
	"encoding/json"
	"fmt"
	"log"
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
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	log.Println(req.ParseForm())
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		Header: req.Header,
	}
}

func (c *Context) PostForm(k string) string{
	// c.Req.FormValue 也能取到 post 请求的 ?username=abc&password=123
	// c.Req.PostFormValue 只会取到 -d "username=abcde&password=12345"
	return c.Req.PostFormValue(k)
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
