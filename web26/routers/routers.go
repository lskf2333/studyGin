package routers

import (
	"github.com/gin-gonic/gin"
	"studyGin/web26/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//告诉gin框架模板文件的静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架那里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateTodo)
		//查看
		//查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdataTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	return r
}
