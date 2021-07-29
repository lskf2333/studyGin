package main

//gin 初使用

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message":"hello golang",
	})
}

func main()  {
	r:=gin.Default() //返回默认的路由引擎
	r.GET("/hello",sayHello)
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"DELETE",
		})
	})

	//启动服务
	r.Run(":9090")
}
