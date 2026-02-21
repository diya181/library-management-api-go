package main

import (
	"github.com/gin-gonic/gin"
	"library-api/database"
	"library-api/handlers"
)

func main() {

	database.Connect()

	r := gin.Default()

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Library API Running",
		})
	})

	// Book Routes
	r.POST("/books", handlers.AddBook)
	r.GET("/books", handlers.GetBooks)
	r.POST("/checkout", handlers.CheckoutBook)
	r.POST("/return", handlers.ReturnBook)
	r.POST("/reserve", handlers.ReserveBook)

	r.Run(":8080")
}