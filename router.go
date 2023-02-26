package main

import (
	"strconv"

	"github.com/gin-gonic/gin"

	_ "todolist/model"
	"todolist/service"

)

// Add todo item godoc
// @Summary Add todo item
// @Schemes
// @Description adds todolist item
// @Tags TodoItem
// @Param todoItem body model.TodoItem true "todo item for add"
// @Accept json
// @Produce json
// @Success 200
// @Router /add [post]
func AddTodoItem(ctx *gin.Context) {
	var (
		name        = ctx.Param("name")
		description = ctx.Param("description")
	)

	todoItemId := service.AddTodoItem(name, description)
	ctx.Query(strconv.Itoa(todoItemId))
}

// Update todo item godoc
// @Summary Update todo item
// @Schemes
// @Description updates todolist item
// @Tags TodoItem
// @Param todoItem body model.TodoItem true "todo item for update"
// @Accept json
// @Produce json
// @Success 200
// @Router /update [put]
func UpdateTodoItem(ctx *gin.Context) {
	var (
		id, _       = strconv.Atoi(ctx.Param("id"))
		name        = ctx.Param("name")
		description = ctx.Param("description")
	)

	service.UpdateTodoItem(id, name, description)
}

// Get todo item by id godoc
// @Summary Get todo item by id
// @Schemes
// @Description gets todolist item
// @Tags TodoItem
// @Param id path int true "id for search todo item"
// @Accept json
// @Produce json
// @Success 200 {object} model.TodoItem
// @Router /get [get]
func GetTodoItemById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	todoItem := service.GetTodoItemById(id)
	ctx.JSON(200, todoItem)
}
