package rest

import (
	"todo-list-api/backend/internal/transport/rest/todoController"
	"todo-list-api/backend/internal/transport/rest/userController"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("api/todos", todoController.GetTodos)
	router.GET("api/todo/:id", todoController.GetTodo)
	router.POST("api/todo", todoController.CreateTodo)
	router.PUT("api/todo/:id", todoController.UpdateTodo)
	router.DELETE("api/todo", todoController.DeleteTodo)

	router.GET("/api/user/:id", userController.GetUser)
	router.PUT("/api/user/:id", userController.UpdateUser)
	router.POST("/api/signup", userController.SignUp)
	router.POST("api/login", userController.Login)

	return router
}
