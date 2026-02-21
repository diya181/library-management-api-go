package handlers

import (
	"net/http"

	"library-api/database"
	"library-api/models"

	"github.com/gin-gonic/gin"
)

func AddBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.AvailableCopies = book.TotalCopies

	database.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{
		"message": "Book added successfully",
		"book":    book,
	})
}

func GetBooks(c *gin.Context) {
	var books []models.Book

	database.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}