package controllers

import (
	"net/http"

	"github.com/FelipeMP0/go-rest-api-gin/v2/models"
	"github.com/gin-gonic/gin"
)

// CreateBookInput is the schema used the validate the creation of new books.
type CreateBookInput struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// GetBooks returns all books.
// GET /books
func GetBooks(c *gin.Context) {
	var books []models.Book
	
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook creates a new book
// POST /books
func CreateBook(c *gin.Context) {
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Author: input.Author}

	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
