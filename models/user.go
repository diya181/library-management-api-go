package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Role string // student or librarian
}