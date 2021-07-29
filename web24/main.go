package main

//gorm  删除

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

	//5. 删除
	//如果字段中有  delete_at 就会软删除，如果没有的话，就会物理删除
	//var u =User{}
	//u.ID=1  //如果没有主键或者主键为空的话，会把所有的数据都删除了
	//db.Debug().Delete(&u)

	//db.Debug().Where("name=?","不存在").Delete(&User{})

	//var u1 []User
	//db.Debug().Unscoped().Find(&u1)   //Unscoped()  可以查到软删除的值
	//fmt.Printf("[]user:%#v\n",u1)

	//物理删除
	//db.Unscoped().Where("name=?","不存在").Delete(&User{})

}
