package controllers

import (
	"net/http"
	"strconv"

	"mytest/models"

	"github.com/gin-gonic/gin"
)

var notes = []models.Note{}
var nextID = 1

func GetNotes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, notes)
}

func CreateNote(ctx *gin.Context) {
	var newNote models.Note
	if err := ctx.ShouldBindJSON(&newNote); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newNote.ID = nextID
	nextID++
	notes = append(notes, newNote)

	ctx.JSON(http.StatusCreated, newNote)
}

func GetNoteById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	for _, note := range notes {
		if (note.ID == id) {
			ctx.JSON(http.StatusOK, note)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})
}

func DeleteNote(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:1], notes[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Note not found"})

}
