package main

import (
	"github.com/zhangzw001/learnGee/3-http-router/gee"
	"net/http"
)

//https://geektutu.com/post/gee-day2.html
func main() {
	e := gee.NewEngine()
	e.GET("/header/:name",headerHandler)
	// Queue 获取get参数
	e.GET("/logout",logoutHandler)
	// POSTForm 获取post参数
	e.POST("/login",loginHandler)
	e.Run("localhost:8000")
}

func headerHandler(ctx *gee.Context) {
	for k,v := range ctx.Header {
		if ctx.Param("name") == k {
			ctx.String(200,"Get header[%s] = %s\n",k,v)
			break
		}
	}
}


func loginHandler(ctx *gee.Context) {

	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	ctx.JSON(http.StatusOK,gee.H{
		"username": username,
		"password": password,
	})
}

func logoutHandler(ctx *gee.Context) {
	username := ctx.Query("username")
	ctx.JSON(http.StatusOK,gee.H{
		"username":username,
	})
}

