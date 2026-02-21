package models

import (
	"time"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID          uint
	BookID          uint
	ReservationDate time.Time
	Status          string
}