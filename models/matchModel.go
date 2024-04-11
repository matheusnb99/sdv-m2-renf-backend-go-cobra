package models

import "gorm.io/gorm"

type Match struct {
	gorm.Model
	MatchDate string `gorm:"not null"`
	HomeTeamID uint `gorm:"not null"`
	HomeTeam Team `gorm:"foreignKey:HomeTeamID"`
	AwayTeamID uint `gorm:"not null"`
	AwayTeam Team `gorm:"foreignKey:AwayTeamID"`
	HomeTeamScore uint
	AwayTeamScore uint 
	TournamentID uint 
	Tournament Tournament `gorm:"foreignKey:TournamentID"`
	WinnerID uint 
	Winner Team `gorm:"foreignKey:WinnerID"`
}

