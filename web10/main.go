package main

//gin 放回json 数据

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//方法一： 使用map
	r.GET("/json", func(c *gin.Context) {
		//data:=map[string]interface{}{
		//	"name":"xixi",
		//	"message":"hello world!",
		//	"age":18,
		//}
		data := gin.H{"name": "xixi", "message": "hello world!", "age": 18}
		c.JSON(http.StatusOK, data)
	})

	//方法二 :结构体
	type msg struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Age     int    `json:"age"`
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			Name:    "haha",
			Message: "hello world!",
			Age:     19,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":9090")
}
