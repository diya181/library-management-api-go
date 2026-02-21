package database

import (
	"fmt"
	"library-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=DDP2123@ dbname=library port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	fmt.Println("Database connected successfully")

	DB = db

	err = DB.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Checkout{},
		&models.Reservation{},
	)

	if err != nil {
		panic("Migration failed")
	}
}