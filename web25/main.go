package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

//Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func (Todo) TableName() string {
	return "todo"
}

func initMySQL() (err error) {
	dsn := "root:root@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	//创建数据库
	//sql:create database bubble
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close() //程序退出关闭数据库
	//模型绑定
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	//告诉gin框架模板文件的静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架那里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v1
	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			if err = DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//查看
		//查看所有的待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "无效的id"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, todo)
		})
		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "无效的id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{id: "deleted"})
		})
	}

	r.Run(":9090")
}
