package models

import (
	"gorm.io/gorm"
)

type Balance struct {
	gorm.Model
	UserID   uint   `gorm:"UNIQUE_INDEX:compositeindex;index;not null"`
	Currency string `gorm:"UNIQUE_INDEX:compositeindex;type:varchar(3);not null"`
	Value    float64
}
