package main

import (
	"log"
	apod "nied-science/internal/api"
	"nied-science/internal/auth"
	"nied-science/internal/repository"
	"nied-science/internal/service"

	db "nied-science/internal/database"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init()

	repo := repository.NewAPODRepository()
	apodService := service.NewAPODService(repo)
	service.StartCronJob(apodService)

	router := gin.Default()
	router.Use(auth.APIKeyMiddleware())

	// Middleware de validação da API Key

	apod.RegisterRoutes(router, apodService)
	if err := router.Run(":8079"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
