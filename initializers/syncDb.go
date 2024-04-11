package initializers

import "github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{}, &models.Sport{}, &models.Team{}, &models.Tournament{}, &models.Match{})
}