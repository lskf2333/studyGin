package main

//获取form 表单提交的参数

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		//username:=c.PostForm("username")
		//password:=c.PostForm("password")

		//username:=c.DefaultPostForm("username","somebody")
		//password:=c.DefaultPostForm("password","88888888")

		username, ok := c.GetPostForm("username")
		if !ok {
			username = "大帅比"
		}
		password, ok := c.GetPostForm("password")
		if !ok {
			password = "******"
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Username": username,
			"Password": password,
		})
	})
	r.Run(":9090")
}
