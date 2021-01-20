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
	Req    *http.Request

	// request info
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int

	// middleware
	handlers []HandlerFunc
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index: -1,
	}
}

func (c *Context) Next() {
	c.index++
	//index是记录当前执行到第几个中间件，当在中间件中调用Next方法时，控制权交给了下一个中间件，直到调用到最后一个中间件，然后再从后往前，调用每个中间件在Next方法之后定义的部分
	for c.index < len(c.handlers) {
		// 这里依次执行 中间件 的方法
		// 注意中间件中必须写c.Next(), 否则后续中间件无法继续执行
		c.handlers[c.index](c)
		c.index++
	}
	/*
	func A(c *Context) {
	    part1
	    c.Next()
	    part2
	}
	func B(c *Context) {
	    part3
	    c.Next()
	    part4
	}
	假设我们应用了中间件 A 和 B，和路由映射的 Handler。c.handlers是这样的[A, B, Handler]，c.index初始化为-1。调用c.Next()，接下来的流程是这样的：

	1 c.index++，c.index 变为 0
	2 0 < 3，调用 c.handlers[0]，即 A
	3 执行 part1，调用 c.Next()
	4 c.index++，c.index 变为 1
	5 1 < 3，调用 c.handlers[1]，即 B
	6 执行 part3，调用 c.Next()
	7 c.index++，c.index 变为 2
	8 2 < 3，调用 c.handlers[2]，即Handler
	9 Handler 调用完毕，返回到 B 中的 part4，执行 part4
	10 part4 执行完毕，返回到 A 中的 part2，执行 part2
	11 part2 执行完毕，结束。
	最终的顺序是part1 -> part3 -> Handler -> part 4 -> part2
	 */
}


func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) PostForm(k string) string {
	return c.Req.FormValue(k)
}

func (c *Context) Query(k string) string {
	return c.Req.URL.Query().Get(k)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(k string, v string) {
	c.Writer.Header().Set(k, v)
}

func (c *Context) String(code int, format string, value ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, value...)))
}

func (c *Context) JSON(code int, obj interface{}) {
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
