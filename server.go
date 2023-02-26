package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"todolist/common"
	"todolist/docs"

)

func main() {
	db := common.Init()
	defer db.Close()

	r := gin.Default()
	addRoutes(r)
	addSwagger(r)

	r.Run()
}

func addRoutes(r *gin.Engine) {
	apiGroup := r.Group("/api")
	todolistGroup := apiGroup.Group("/todolist")

	todolistGroup.POST("/add", AddTodoItem)
	todolistGroup.PUT("/update", UpdateTodoItem)
	todolistGroup.GET("/get", GetTodoItemById)
}

func addSwagger(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/todolist"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
