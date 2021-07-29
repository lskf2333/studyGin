package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1. 定义模型
type User struct {
	ID   int64
	Name sql.NullString `gorm:"default:'haha'"` //当字段设有默认值的时候，有需要保存对应类型的0值，可以使用sql.NullType获取使用指针
	Age  *int64         `gorm:"default:18"`
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3. 创建记录
	u := User{Name: sql.NullString{"", true}, Age: new(int64)} //在代码中创建一个user对象
	fmt.Println(db.NewRecord(&u))                              //判断主键是否为空  true
	db.Create(&u)
	fmt.Println(db.NewRecord(&u)) //判断主键是否为空   false
}
