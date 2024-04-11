package controllers

import (
	"fmt"
	"strconv"

	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/initializers"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/models"
)

type MatchDetail struct {
	ID         uint
	ScoreTeamAway	string
	ScoreTeamLocal	string
	Win        bool
	MatchDate  string
}



func CreateMatch(HomeTeamID uint, AwayTeamID uint, MatchDate string) string {
	if HomeTeamID == 0 {
		panic("HomeTeamID is required")
	}

	if AwayTeamID == 0 {
		panic("AwayTeamID is required")
	}

	if MatchDate == "" {
		panic("MatchDate is required")
	}

	match := models.Match{HomeTeamID: HomeTeamID, AwayTeamID: AwayTeamID, MatchDate: MatchDate}
	result := initializers.DB.Create(&match)

	if result.Error != nil {
		panic("error creating match")
	}


	return "Match created successfully"
}



func GetMatches() []models.Match {
	var matchs []models.Match
	result := initializers.DB.Find(&matchs)

	if result.Error != nil {

		panic("error fetching matchs")
	}

	return matchs

}

func GetMatchesByTeamID(teamID uint) []models.Match {
	var matchs []models.Match
	result := initializers.DB.Where("home_team_id = ? OR away_team_id = ?", teamID, teamID).Find(&matchs)

	if result.Error != nil {
		panic("error fetching matchs")
	}

	return matchs
}

func GetMatchesByName(teamName string) []MatchDetail {
	var matches []models.Match
	result := initializers.DB.Joins("JOIN teams ON home_team_id = teams.id OR away_team_id = teams.id").Where("teams.name = ?", teamName).Find(&matches)

	if result.Error != nil {
		panic("error fetching matchs")
	}

	var matchDetailsList []MatchDetail


	for _, match := range matches {
		matchDetails, err := GetMatchDetails(match.ID)
		if err != nil {
			panic("error fetching match details")
		}
		matchDetailsList = append(matchDetailsList, matchDetails)
	}
	return matchDetailsList

}


func GetMatchDetails(matchId uint) (MatchDetail, error) {
	var match models.Match
	result := initializers.DB.Preload("HomeTeam").Preload("AwayTeam").First(&match, matchId).Where("matches.id = ?", matchId).First(&match)

	if result.Error != nil {
		fmt.Println("error fetching match details")
	}

	win := match.WinnerID == match.HomeTeamID || match.WinnerID == match.AwayTeamID
	ScoreTeamAway := match.AwayTeam.Name +" " +  strconv.Itoa(int(match.AwayTeamScore))
	ScoreTeamLocal := strconv.Itoa(int(match.HomeTeamScore)) +" " + match.HomeTeam.Name


	return MatchDetail{
		ID:         match.ID,
		ScoreTeamAway:	ScoreTeamAway,
		ScoreTeamLocal:	ScoreTeamLocal,
		Win:        win,
		MatchDate:  match.MatchDate,
	}, nil

}

func GetMatchWithTeams() ([]MatchDetail, error) {
	matches := GetMatches()
	
	var matchDetailsList []MatchDetail

	for _, match := range matches {
		matchDetails, err := GetMatchDetails(match.ID)
		if err != nil {
			panic("error fetching match details")
		}
		matchDetailsList = append(matchDetailsList, matchDetails)
	}
	return matchDetailsList, nil

}


func MatchDetailToString(match MatchDetail) string {
	return fmt.Sprintf("id: %d, %s, %s, win: %t, matchDate: %s", match.ID,  match.ScoreTeamLocal,match.ScoreTeamAway, match.Win, match.MatchDate)
}

func DeleteMatch(matchId uint) string {
	var match models.Match
	result := initializers.DB.Delete(&match, matchId)

	if result.Error != nil {
		panic("error deleting match")
	}

	return "Match deleted successfully"
}



func GetMatchTeamsNames(matchId uint) (string, string) {
		var match models.Match
		result := initializers.DB.Preload("HomeTeam").Preload("AwayTeam").First(&match, matchId).Where("matches.id = ?", matchId).First(&match)
		
		if result.Error != nil {
			fmt.Println("error fetching match details")
			return "", ""
		}

	return match.HomeTeam.Name, match.AwayTeam.Name
}

func ChangeScoreOfTeam(matchId uint, teamId uint, newScore uint) string {
	var match models.Match
	result := initializers.DB.First(&match, matchId)

	if result.Error != nil {
		panic("error fetching match")
	}

	if match.HomeTeamID == teamId {
		match.HomeTeamScore = newScore
	} else {
		match.AwayTeamScore = newScore
	}

	result = initializers.DB.Save(&match)

	if result.Error != nil {
		panic("error updating match")
	}

	return "Match updated successfully"
}

func SetMatchOver(matchId uint) string {
	var match models.Match
	result := initializers.DB.First(&match, matchId)

	if result.Error != nil {
		panic("error fetching match")
	}

	// winner is the team with the highest score
	if match.HomeTeamScore > match.AwayTeamScore {
		match.WinnerID = match.HomeTeamID
	} else {
		match.WinnerID = match.AwayTeamID
	}


	result = initializers.DB.Save(&match)

	if result.Error != nil {
		panic("error updating match")
	}

	return "Match updated successfully"
}