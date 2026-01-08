package main

import (
	"log"
	"os"
	"todo-list-api/backend/internal/models"
	"todo-list-api/backend/internal/transport/rest"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env dari root project (2 level di atas dari cmd/app)
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	log.Println("TOKEN:", os.Getenv("TOKEN"))
	log.Println("DB:", os.Getenv("DB"))

	models.ConnectDatabase()

	router := rest.SetupRouter()

	router.Run(":8080")
}
