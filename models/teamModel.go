package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	SportID uint `gorm:"not null"`
	Sport Sport `gorm:"foreignKey:SportID"`

}