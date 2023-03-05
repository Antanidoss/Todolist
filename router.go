package main

import (
	"strconv"

	"github.com/gin-gonic/gin"

	_ "todolist/models"
	"todolist/services"
)

type addTodoItemDto struct {
	Name        string
	Description string
}

// Add todo item godoc
// @Summary Add todo item
// @Schemes
// @Description adds todolist item
// @Tags TodoItem
// @Param todoItem body addTodoItemDto true "todo item for add"
// @Accept json
// @Produce json
// @Success 200
// @Router /add [post]
func AddTodoItem(ctx *gin.Context) {
	todoItem := addTodoItemDto{}
	ctx.ShouldBind(&todoItem)

	todoItemId := services.AddTodoItem(todoItem.Name, todoItem.Description)
	ctx.Query(strconv.Itoa(int(todoItemId)))
}

// Update todo item godoc
// @Summary Update todo item
// @Schemes
// @Description updates todolist item
// @Tags TodoItem
// @Param todoItem body models.TodoItem true "todo item for update"
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

	services.UpdateTodoItem(id, name, description)
}

// Get todo item by id godoc
// @Summary Get todo item by id
// @Schemes
// @Description gets todolist item
// @Tags TodoItem
// @Param id path int true "id for search todo item"
// @Accept json
// @Produce json
// @Success 200 {object} models.TodoItem
// @Router /get/{id} [get]
func GetTodoItemById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		panic("GetTodoItemById:%v" + err.Error())
	}

	todoItem := services.GetTodoItemById(id)
	ctx.JSON(200, todoItem)
}
