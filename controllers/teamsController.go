package controllers

import (
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/initializers"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/models"
)

func CreateTeam(Name string, Sport string) string {
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

	team := models.Team{Name: Name, SportID: sport.ID}
	if err := tx.Create(&team).Error; err != nil {
		tx.Rollback()
		panic("error creating team")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic("error committing transaction")
	}

	return "Team created successfully"
}


func GetTeams() []models.Team {
	var teams []models.Team
	result := initializers.DB.Find(&teams)

	if result.Error != nil {
		panic("error fetching teams")
	}

	return teams
}

func GetTeamNames() []string {
	teams := GetTeams()
	names := make([]string, len(teams))

	for i, team := range teams {
		names[i] = team.Name
	}

	return names
}


func DeleteTeam(Name string) string {
	if Name == "" {
		panic("Name is required")
	}

	tx := initializers.DB.Begin()

	var team models.Team
	if err := tx.First(&team, "name = ?", Name).Error; err != nil {
		tx.Rollback()
		panic("team not found")
	}

	if err := tx.Delete(&team).Error; err != nil {
		tx.Rollback()
		panic("error deleting team")
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic("error committing transaction")
	}

	return "Team deleted successfully"
}