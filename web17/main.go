package main

//gin 路由和路由组

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//访问/index的GET请求会周这一条处理逻辑
	//路由
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mothed": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mothed": "POST",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mothed": "PUT",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mothed": "DELETE",
		})
	})
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"mothed": http.MethodGet})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"mothed": http.MethodPost})
		case http.MethodPut:
			c.JSON(http.StatusOK, gin.H{"mothed": http.MethodPut})
		case http.MethodDelete:
			c.JSON(http.StatusOK, gin.H{"mothed": http.MethodDelete})
		}
	})

	//视频的首页和详情页
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{"msg":"/video/index"})
	//})
	//r.GET("/video/detail", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{"msg":"/video/detail"})
	//})
	//路由组 多用于区分不同的业务线或api版本
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})
		videoGroup.GET("/detail", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/detail"})
		})
	}

	//商城的首页和详情页
	//r.GET("/shop/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{"msg":"/shop/index"})
	//})
	//r.GET("/shop/detail", func(c *gin.Context) {
	//	c.JSON(http.StatusOK,gin.H{"msg":"/shop/detail"})
	//})
	//路由组
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
		})
		shopGroup.GET("/detail", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/shop/detail"})
		})
		//路由组的嵌套
		xixi := shopGroup.Group("xixi")
		xixi.GET("/haha", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/shop/xixi/haha"})
		})
	}

	//NoRoute
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "没有这个页面",
		})
	})
	r.Run(":9090")
}
