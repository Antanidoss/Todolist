package services

import (
	"fmt"
	"strconv"
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
	sql := fmt.Sprintf(`
		SELECT [Name], [Description]
  		FROM [dbo].[TodolistItems]
  		where ID = '%s'`, strconv.Itoa(id))

	var (
		name        string
		description string
	)
	err := common.DB.QueryRow(sql).Scan(&name, &description)

	if err != nil {
		panic("GetTodoItemById:%v" + err.Error())
	}

	var todoItem models.TodoItem
	todoItem.ID = uint(id)
	todoItem.Name = name
	todoItem.Description = description

	return &todoItem
}
