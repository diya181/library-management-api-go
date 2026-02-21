package handlers

import (
	"net/http"
	"time"

	"library-api/database"
	"library-api/models"

	"github.com/gin-gonic/gin"
)

type CheckoutRequest struct {
	UserID uint `json:"user_id"`
	BookID uint `json:"book_id"`
}

func CheckoutBook(c *gin.Context) {
	var request CheckoutRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book

	// Find book
	if err := database.DB.First(&book, request.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Check availability
	if book.AvailableCopies <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No copies available"})
		return
	}

	// Start transaction
	tx := database.DB.Begin()

	// Reduce available copies
	book.AvailableCopies -= 1
	tx.Save(&book)

	// Create checkout record
	checkout := models.Checkout{
		UserID:       request.UserID,
		BookID:       request.BookID,
		CheckoutDate: time.Now(),
		DueDate:      time.Now().AddDate(0, 0, 7), // 7 days
		FineAmount:   0,
	}

	tx.Create(&checkout)

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message":  "Book checked out successfully",
		"checkout": checkout,
	})
}
func ReturnBook(c *gin.Context) {
	type ReturnRequest struct {
		CheckoutID uint `json:"checkout_id"`
	}

	var request ReturnRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var checkout models.Checkout

	// Find checkout record
	if err := database.DB.First(&checkout, request.CheckoutID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checkout record not found"})
		return
	}

	// Prevent double return
	if checkout.ReturnDate != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book already returned"})
		return
	}

	// Start transaction
	tx := database.DB.Begin()

	now := time.Now()
	checkout.ReturnDate = &now

	// Fine calculation
	if now.After(checkout.DueDate) {
		daysLate := int(now.Sub(checkout.DueDate).Hours() / 24)
		checkout.FineAmount = float64(daysLate * 10) // ₹10 per day
	}

	tx.Save(&checkout)

	// Increase available copies
	var book models.Book
	tx.First(&book, checkout.BookID)
	book.AvailableCopies += 1
	tx.Save(&book)

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message": "Book returned successfully",
		"fine":    checkout.FineAmount,
	})
}