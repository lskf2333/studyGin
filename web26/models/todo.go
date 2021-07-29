package models

import (
	"studyGin/web26/dao"
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

//todo 增删改查

//CreateATodo
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetATodoById(id string) (todo *Todo, err error) {
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodoById(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
