package db

import (
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	Email     string `gorm:"index"` //uniqe rule is set in Init()
	City      string `gorm:"not null"`
	Frequency string `gorm:"not null"`
	Confirmed bool
	Token     string `gorm:"uniqueIndex"`
}
