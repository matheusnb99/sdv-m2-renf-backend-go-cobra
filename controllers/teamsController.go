package controllers

import (
	"errors"
	"fmt"
	"os"

	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/initializers"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/models"
	"gorm.io/gorm"
)

func CreateTeam(Name string, Sport string) string {
	if Name == "" {
		panic("Name is required")
	}

	if Sport == "" {
		panic("Sport is required")
	}

	tx := initializers.DB.Begin()

	fmt.Printf("A \n\n")
	var sport models.Sport
	if err := tx.First(&sport, "name = ?", Sport).Error; err != nil {
		// Check if the error is due to record not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle "record not found" error gracefully
			// For example, you can create a new record here
			sport = models.Sport{Name: Name}
			if err := tx.Create(&sport).Error; err != nil {
				tx.Rollback()
				fmt.Printf("Prompt failed %v\n", err) // Print error for debugging purposes
				os.Exit(1)
			}
		} else {
			// Handle other errors
			tx.Rollback()
			os.Exit(1)
		}
	}
	
	
	
	
	fmt.Printf("B \n\n")
	team := models.Team{Name: Name, SportID: sport.ID}
	if err := tx.Create(&team).Error; err != nil {
		tx.Rollback()
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
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


func DeleteTeam(Name string) (string, error) {
if Name == "" {
	return "", errors.New("name is required")
}

tx := initializers.DB.Begin()

var team models.Team
if err := tx.First(&team, "name = ?", Name).Error; err != nil {
	tx.Rollback()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "Team not found", nil
	}
	return "", fmt.Errorf("error finding team: %w", err)
}

    if err := tx.Delete(&team).Error; err != nil {
        tx.Rollback()
        return "", fmt.Errorf("error deleting team: %w", err)
    }

    if err := tx.Commit().Error; err != nil {
        tx.Rollback()
        return "", fmt.Errorf("error committing transaction: %w", err)
    }

    return "Team deleted successfully", nil
}


func GetIdByName(Name string) (uint, error) {
	if Name == "" {
		return 0, errors.New("name is required")
	}

	var team models.Team
	if err := initializers.DB.First(&team, "name = ?", Name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, fmt.Errorf("error finding team: %w", err)
	}

	return team.ID, nil
}