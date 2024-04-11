package models

import "gorm.io/gorm"

type Tournament struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	SportID uint `gorm:"not null"`
	Sport Sport `gorm:"foreignKey:SportID"`
	Matches []Match
}