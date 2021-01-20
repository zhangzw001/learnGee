package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func main() {
	e := gin.Default()

	e.LoadHTMLGlob("example/upload/template/upload.html")
	e.GET("/upload",uploadHandler)
	e.POST("/uploadFile",uploadFileHandler)

	e.Run("localhost:8888")
}



func uploadHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html",nil)
}

func uploadFileHandler(c *gin.Context) {
	f , err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusBadGateway,gin.H{
			"code":200,
			"msg":"ok",
			"error":err,
		})
		c.Abort()
		return
	}

	c.SaveUploadedFile(f,"/Users/zhangzw/lrzsz/logs/"+f.Filename)
	//src,_ := f.Open()
	//defer src.Close()
	//
	//dst,_ := os.Create("/Users/zhangzw/lrzsz/logs/"+f.Filename)
	//defer dst.Close()
	io.Copy(dst,src)

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"ok",
		"error":"",
	})
	return
}
