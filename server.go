package main

import (
	"github.com/gin-gonic/gin"

	"todolist/common"
)

func main() {
	db := common.Init()
	defer db.Close()

	r := gin.Default()
	addRoutes(r)

	r.Run()
}

func addRoutes(r *gin.Engine) {
	apiGroup := r.Group("api")
	todolistGroup := apiGroup.Group("todolist")

	todolistGroup.POST("/add", AddTodoItem)
	todolistGroup.PUT("/update", UpdateTodoItem)
	todolistGroup.GET("/get", GetTodoItemById)
}
