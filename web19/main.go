package main

//GORM 连接mysql基本实例

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//创建表 自动迁移 （把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	//u1 :=UserInfo{ID: 1,Name: "xixi",Gender: "男",Hobby: "蛙泳"}
	//db.Create(&u1)

	var u UserInfo
	//查询
	db.First(&u) //查询表中第一跳数据保存到u中
	fmt.Printf("u:%#v/n", u)
	//更新
	db.Model(&u).Update("hobby", "双色球")
	//删除
	db.Delete(&u)
}
