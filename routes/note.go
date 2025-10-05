package routes

import (
	"mytest/controllers"

	"github.com/gin-gonic/gin"
)

func NoteRoutes(route *gin.Engine) {
	noteGroup := route.Group("/notes")
	{
		noteGroup.GET("/", controllers.GetNotes)
		noteGroup.POST("/", controllers.CreateNote)
		noteGroup.GET("/:id", controllers.GetNoteById)
		noteGroup.DELETE("/:id", controllers.DeleteNote)
	}
}
