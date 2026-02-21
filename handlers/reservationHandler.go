package handlers

import (
	"net/http"
	"time"

	"library-api/database"
	"library-api/models"

	"github.com/gin-gonic/gin"
)

type ReservationRequest struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}

func ReserveBook(c *gin.Context) {
	var request ReservationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book

	if err := database.DB.First(&book, request.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if book.AvailableCopies > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Book is available. You can checkout directly.",
		})
		return
	}

	reservation := models.Reservation{
		UserID:          request.UserID,
		BookID:          request.BookID,
		ReservationDate: time.Now(),
		Status:          "waiting",
	}

	database.DB.Create(&reservation)

	c.JSON(http.StatusOK, gin.H{
		"message":     "Reservation added to waiting list",
		"reservation": reservation,
	})
}