package main

import (
	"mytest/config"
	"mytest/models"
	"mytest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect DB
	config.ConnectDatabase()

	// Migrate models
	config.DB.AutoMigrate(&models.Note{})

	// Setup routes
	routes.NoteRoutes(router)

	router.Run("localhost:8080")
}
