package main

import (
	"todo-list-api/backend/internal/models"
	"todo-list-api/backend/internal/transport/rest"
)

func main() {
	models.ConnectDatabase()

	r := rest.SetupRouter()
	r.Run()
}
