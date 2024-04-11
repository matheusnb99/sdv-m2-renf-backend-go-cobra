package controllers

import (
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/initializers"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/models"
)


func CreateSport(Name string) string {

	if Name == "" {
		panic("Name is required")
	}

	sport := models.Sport{Name: Name}
	result := initializers.DB.Create(&sport)

	if result.Error != nil {
		panic("error creating sport")
	}

	return "Sport created successfully"
}

func GetSports() []models.Sport {
	var sports []models.Sport
	result := initializers.DB.Find(&sports)

	if result.Error != nil {
		panic("error fetching sports")
	}

	return sports
}


func GetSportNames() []string {
	sports := GetSports()
	names := make([]string, len(sports))

	for i, sport := range sports {
		names[i] = sport.Name
	}

	return names
}


func CreateSportReturning(Name string) models.Sport {

	if Name == "" {
		panic("Name is required")
	}

	sport := models.Sport{Name: Name}
	result := initializers.DB.Create(&sport)

	if result.Error != nil {
		panic("error creating sport")
	}

	return sport
}

