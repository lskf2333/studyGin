package main

//gin 请求重定向

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/haha", func(c *gin.Context) {
		//fmt.Println("xixixix")
		//c.JSON(http.StatusOK,gin.H{
		//	"status":"ok",
		//})
		//跳转到sogo.com
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com")
	})
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b" //把请求的URI 修改
		r.HandleContext(c)        //继续后续的处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	r.Run(":9090")
}
