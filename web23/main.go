package main

//gorm 更新

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 1. 创建模型
type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
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

	//4.查询
	var user User
	db.First(&user)

	//5. 更新
	//user.Name="jaja"
	//user.Age=20
	//db.Debug().Save(&user) //默认会修改所有记录
	//db.Debug().Model(&user).Update("name","kaka")

	//修改多个字段
	//m1:=map[string]interface{}{
	//	"name":"gaga",
	//	"age":21,
	//	"active":true,
	//}
	//db.Debug().Model(&user).Updates(m1) //m1列出的字段都会修改
	//db.Debug().Model(&user).Select("age").Updates(m1) //只更新age字段
	//db.Debug().Model(&user).Omit("active").Updates(m1)	//更新m1中除了active 的字段

	//db.Model(&user).UpdateColumn("name", "hello") // UPDATE users SET name='hello' WHERE id = 1;

	//使用sql表达式更新
	//让users表中所有的用户的年龄在原来基础上+2
	db.Debug().Model(&User{}).Update("age", gorm.Expr("age+?", 2)) // UPDATE `users` SET `age` = age+2, `updated_at` = '2021-07-06 16:06:27'  WHERE `users`.`deleted_at` IS NULL

}
