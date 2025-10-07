package controllers

import (
	"net/http"
	"strconv"

	"mytest/config"
	"mytest/models"

	"github.com/gin-gonic/gin"
)

func GetNotes(ctx *gin.Context) {
	var notes []models.Note
	config.DB.Find(&notes)
	ctx.JSON(http.StatusOK, notes)
}

func CreateNote(ctx *gin.Context) {
	var newNote models.Note
	if err := ctx.ShouldBindJSON(&newNote); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&newNote)
	ctx.JSON(http.StatusCreated, newNote)
}

func GetNoteById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var note models.Note
	if err := config.DB.First(&note, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
		return
	}
	ctx.JSON(http.StatusOK, note)
}

func DeleteNote(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var note models.Note
	if err := config.DB.First(&note, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
		return
	}
	config.DB.Unscoped().Delete(&note)
	ctx.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
