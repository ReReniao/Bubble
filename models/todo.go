package models

import (
	"bubble/dao"
)
// Todo Model
type Todo struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

/*
	Todo这个model的增删改查操作都放这里
 */

// Todo增删改查
// CreateATodo 创建Todo
func CreateATodo(todo *Todo) (err error){
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodo(id string) (todo *Todo, err error){
	todo = new(Todo)
	if err = dao.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return 
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&Todo{}).Error
	return
}