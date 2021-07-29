package main

//GORM gorm 模型定义

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"`
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`              // 设置字段大小为255
	MemberNumber *string `gorm:"not null;unique_index"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`        // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`            // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`                     // 忽略本字段
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

func (Animal) TableName() string {
	return "animal"
}

type Test struct {
	Id   int64 `gorm:"primary_key"`
	Name string
}

func main() {

	//设置表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "gorm_" + defaultTableName
	}

	//连接MySQL数据库
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true) //禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	//创建表 自动迁移 （把结构体和数据表进行对应）
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	//创建表名
	//db.Table("test").CreateTable(&Test{})
}
