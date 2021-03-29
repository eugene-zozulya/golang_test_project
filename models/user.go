package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Balance     []Balance
	Transaction []Transaction
}
