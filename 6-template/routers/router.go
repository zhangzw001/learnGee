package routers

import (
	"fmt"
	"github.com/zhangzw001/learnGee/6-template/gee"
	"github.com/zhangzw001/learnGee/6-template/middleware"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d ", year, month, day,t.Hour(),t.Minute(),t.Second())
}

func InitRouter() *gee.Engine {
	e := gee.NewEngine()
	e.Use(middleware.Logger())
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

	// static 测试
	//e.Static("/assets","/Users/zhangzw/oschina/golang/golanglearn/src/github.com/zhangzw001/learnGee/assets")
	// template
	e.SetFuncMap(template.FuncMap{
		"FormatAsDate":FormatAsDate,
	})
	e.LoadHTMLGlob("6-template/templates/*")
	e.Static("/assets","6-template/static")
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	e.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	e.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	e.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			//"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
			"now":   time.Now(),
			"now2":   time.Now(),
		})
	})

	// panic
	e.GET("/panic",func(c *gee.Context) {
		names := []string{"test"}
		c.String(http.StatusOK,names[1])
	})
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

