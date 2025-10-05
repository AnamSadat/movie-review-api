package main

import (
	"mytest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.NoteRoutes(router)
	router.Run("localhost:8080")
}