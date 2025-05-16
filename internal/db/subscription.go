package db

import (
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex"`
	City      string `gorm:"not null"`
	Frequency string `gorm:"not null"`
	Confirmed bool
	Token string `gorm:"uniqueIndex"`
}
