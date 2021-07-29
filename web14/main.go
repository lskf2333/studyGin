package main

//gin参数绑定

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		//username:=c.Query("username")
		//password:=c.Query("password")
		//u:=UserInfo{
		//	Username: username,
		//	Password: password,
		//}
		var u UserInfo //声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			fmt.Printf("%#v", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.GET("/form", func(c *gin.Context) {
		var u UserInfo //声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			fmt.Printf("%#v", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.Run(":9090")
}
