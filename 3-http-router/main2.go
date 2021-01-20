package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
)


func main() {
	e := gin.Default()
	log.Println(int64(10 << 20),int64(1<<63 - 1))
	e.Any("/test",func(c *gin.Context) {
		// 获取所有参数
		_ = c.Request.ParseForm()
		log.Println(c.Request.Header["Content-Type"])
		// get form 参数, c.Request.URL.Query, err  = url.ParseQuery
		log.Println(c.Request.URL.Query())
		//log.Println(url.ParseQuery(c.Request.URL.RawQuery))
		// post form 参数
		log.Println(c.Request.PostForm)
		// form 参数
		log.Println(c.Request.Form)
		value ,_ := url.ParseQuery("abc=123&q=123&ccc=123123")
		log.Println(value)

		//m := make(url.Values)
	})

	e.Run("localhost:8123")
}
