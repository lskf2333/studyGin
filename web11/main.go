package main

//gin 获取请求路径上的参数

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的query_string 参数
		name := c.Query("query") //通过query获取请求中携带的参数
		age := c.Query("age")    //通过query获取请求中携带的参数

		//name:=c.DefaultQuery("query","somebody") //取不到就用指定的默认值

		//name ,ok :=c.GetQuery("query") //取不到  第二个参数返回false
		//if !ok {
		//	name ="somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9090")
}
