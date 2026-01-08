package rest

import (
	"todo-list-api/backend/internal/transport/rest/middleware"
	"todo-list-api/backend/internal/transport/rest/todoController"
	"todo-list-api/backend/internal/transport/rest/userController"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500", "http://localhost:5500", "http://127.0.0.1:5501", "http://localhost:5501"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
		AllowCredentials: true,
	}))

	//public route
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.GET("/validate", middleware.RequireAuth, userController.Validate)

	//require auth
	protected := router.Group("/")
	protected.Use(middleware.RequireAuth)
	{
		//todo routes
		protected.GET("/todos", todoController.GetTodos)
		protected.GET("/todos/:id", todoController.GetTodo)
		protected.POST("/todos", todoController.CreateTodo)
		protected.PUT("/todos/:id", todoController.UpdateTodo)
		protected.DELETE("/todos/:id", todoController.DeleteTodo)
		protected.PATCH("/todos/toggle", todoController.ToggleTodo)

		//user routes
		protected.GET("/user/:id", userController.GetUser)
		protected.PUT("/user/:id", userController.UpdateUser)
		protected.POST("/logout", userController.Logout)
	}
	return router
}
