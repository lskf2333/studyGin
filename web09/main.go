package main

//gin 渲染模板文件  和调用静态文件

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/xxx", "./statics")
	//gin自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//解析模板
	//r.LoadHTMLFiles("./templates/posts/index.tmpl","./templates/users/index.tmpl")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		//渲染模板
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "<a href='http://www.baidu.com'>baidu</a>",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		//渲染模板
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "users/index",
		})
	})
	r.Run(":9090")
}
