package controllers

import (
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/initializers"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/models"
)

func CreateTournament(Name string, Sport string) string {
	if Name == "" {
		panic("Name is required")
	}

	if Sport == "" {
		panic("Sport is required")
	}

	tx := initializers.DB.Begin()

	var sport models.Sport
	if err := tx.First(&sport, "name = ?", Sport).Error; err != nil {
		sport = models.Sport{Name: Name}
		if err := tx.Create(&sport).Error; err != nil {
			tx.Rollback()
			panic("error creating sport")
		}
	}

	tournament := models.Tournament{Name: Name, SportID: sport.ID}
	if err := tx.Create(&tournament).Error; err != nil {
		tx.Rollback()
		panic("error creating tournament")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic("error committing transaction")
	}

	return "Tournament created successfully"
}
