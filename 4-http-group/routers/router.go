package routers

import (
	"github.com/zhangzw001/learnGee/4-http-group/gee"
	"github.com/zhangzw001/learnGee/4-http-group/middleware"
	"net/http"
)

func InitRouter() *gee.Engine {
	e := gee.NewEngine()
	e.Use(middleware.Logger(),middleware.Logger2())
	e.GET("/index",indexHandler)

	// v1 group
	groupV1 := e.Group("/v1")
	{
		groupV1.GET("/header/:name", headerHandler)
	}

	// admin
	groupAdmin := e.Group("/admin")
	groupAdmin.Use(middleware.OnlyForV2())
	{
		groupAdminV1:=groupAdmin.Group("/v1")
		{
			// Queue 获取get参数
			groupAdminV1.GET("/logout", logoutHandler)
			// POSTForm 获取post参数
			groupAdminV1.POST("/login", loginHandler)
		}

	}

	return e
}



func headerHandler(ctx *gee.Context) {
	for k,v := range ctx.Req.Header {
		if ctx.Param("name") == k {
			ctx.String(200,"Get header[%s] = %s\n",k,v)
			break
		}
	}
}

func indexHandler(ctx *gee.Context) {
	ctx.JSON(200,gee.H{
		"code":200,
		"msg":"ok",
	})
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

