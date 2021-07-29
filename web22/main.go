package main

//gorm 查询

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1. 创建模型
type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//2.模型和数据表绑定
	db.AutoMigrate(&User{})

	//3.创建记录
	//u1:=User{Name: "xixi",Age: 19}
	//db.Create(&u1)
	//u2:=User{Name: "haha",Age: 18}
	//db.Create(&u2)

	// 4. 查询
	var user User
	//var users []User
	//db.First(&user) // SELECT * FROM users ORDER BY id LIMIT 1;
	//fmt.Printf("user:%#v\n",user)
	//
	//// 随机获取一条记录
	//db.Take(&user) // SELECT * FROM users LIMIT 1;
	//fmt.Printf("user:%#v\n",user)
	//// 根据主键查询最后一条记录
	//db.Last(&user) // SELECT * FROM users ORDER BY id DESC LIMIT 1;
	//fmt.Printf("user:%#v\n",user)
	//// 查询所有的记录
	//db.Find(&users) // SELECT * FROM users;
	//fmt.Printf("user:%#v\n",users)
	//// 查询指定的某条记录(仅当主键为整型时可用)
	//db.First(&user, 10) // SELECT * FROM users WHERE id = 10;
	//fmt.Printf("user:%#v\n",user)

	//一般的查询
	//db.Where("name = ?", "haha").First(&user) // SELECT * FROM users WHERE name = 'haha' limit 1;
	//fmt.Printf("user:%#v\n",user)

	// Struct
	//db.Where(&User{Name: "jinzhu", Age: 20}).First(&user) // SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

	// Map
	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users) // SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	//FirstOrCreate 根据条件获取第一条记录，如果不存在就创建该记录
	db.FirstOrCreate(&user, User{Name: "不存在", Age: 20})

}
