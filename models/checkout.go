package models

import (
	"time"
	"gorm.io/gorm"
)

type Checkout struct {
	gorm.Model
	UserID       uint
	BookID       uint
	CheckoutDate time.Time
	DueDate      time.Time
	ReturnDate   *time.Time
	FineAmount   float64
}