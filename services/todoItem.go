package services

import (
	"fmt"
	"todolist/common"
	"todolist/models"
)

func AddTodoItem(name string, description string) int64 {
	sql := fmt.Sprintf(`
		INSERT INTO [dbo].[TodolistItems]
		([Name], [Description])
		OUTPUT Inserted.ID
		VALUES ('%s','%s')`, name, description)

	var todoItemId int64
	err := common.DB.QueryRow(sql).Scan(&todoItemId)

	if err != nil {
		panic("AddTodoItem:%v" + err.Error())
	}

	return todoItemId
}

func UpdateTodoItem(id int, name string, description string) {

}

func GetTodoItemById(id int) *models.TodoItem {
	return nil
}
