package main

//中间件

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name") //从上下文中取值
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "index",
			"name": name,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "获取不到",
		})
	}

}

//定义一个中间件m1:统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")
	//计时
	start := time.Now()
	//go funcXX(c.Copy())  //在funcXX 中只能使用c的拷贝，这样可以避免在funcXX 中修改了c,导致后续的结果不可控
	c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Set("name", "xixi") //给后面的处理函数传递参数
	c.Next()              //调用后续的处理函数
	//c.Abort() //阻止调用后续的处理函数
	fmt.Println("m2 out...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者一些其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			//是否登录的判断
			//if 是登录用户
			c.Next()
			//else
			//c.Abort
		} else {
			c.Next()
		}

	}
}

func main() {
	r := gin.Default()                  //默认使用了Logger 和 Recovery 的中间件  如果不想使用这两个中间件的话，可以直接使用gin.New()
	r.Use(m1, m2, authMiddleware(true)) //全局注册中间件函数吗
	r.GET("/index", indexHandler)
	r.GET("/detail", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "detail",
		})
	})

	//路由组
	shopGroup := r.Group("/shop")
	shopGroup.Use(authMiddleware(true)) //路由组注册中间件
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "index",
			})
		})
	}

	r.Run(":9090")
}
