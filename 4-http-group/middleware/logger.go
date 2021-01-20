package middleware

import (
	"github.com/zhangzw001/learnGee/4-http-group/gee"
	"log"
	"time"
)

func Logger() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
		c.Next()
	}
}


func Logger2() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		log.Printf("Logger2 Start [%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
		c.Next()
		log.Printf("Logger2 End [%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))

	}
}
func OnlyForV2() gee.HandlerFunc {
	return func(c *gee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
