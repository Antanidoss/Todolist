package main

import (
	"strconv"
	"todolist/service"

	"github.com/gin-gonic/gin"
)

func AddTodoItem(ctx *gin.Context) {
	var (
		name        = ctx.Param("name")
		description = ctx.Param("description")
	)

	todoItemId := service.AddTodoItem(name, description)
	ctx.Query(strconv.Itoa(todoItemId))
}

func UpdateTodoItem(ctx *gin.Context) {
	var (
		id, _       = strconv.Atoi(ctx.Param("id"))
		name        = ctx.Param("name")
		description = ctx.Param("description")
	)

	service.UpdateTodoItem(id, name, description)
}

func GetTodoItemById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	todoItem := service.GetTodoItemById(id)
	ctx.JSON(200, todoItem)
}
