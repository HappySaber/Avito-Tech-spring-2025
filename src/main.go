package main

import (
	"PVZ/src/database"
	routes "PVZ/src/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("src/.env")

	if err != nil {
		log.Fatalf("Error while loading .env file: %v", err)
	}

	database.Init()
	port := "8080"
	router := gin.New()
	routes.PVZRoutes(router)
	router.Run(":" + port)
}
