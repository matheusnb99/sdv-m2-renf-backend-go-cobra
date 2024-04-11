/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package match

import (
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/controllers"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/utils"
	"github.com/spf13/cobra"
)


type promptContent struct {
	errorMsg string
	label    string
}


// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new match.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		createNewTeam()
	},
}

func init() {
	MatchCmd.AddCommand(createCmd)
}

func createNewTeam() {
	teams := controllers.GetTeamNames()
	
	home_team := utils.PromptGetSelect(utils.PromptContent{
			ErrorMsg: "Sport cannot be empty",
			Label:    "Select the sport of the team",
	}, teams)
	home_team_id, err := controllers.GetIdByName(home_team)

	if err != nil {
		panic("error getting id")
	}

	teams = utils.Filter(teams, home_team)
	
	away_team := utils.PromptGetSelect(utils.PromptContent{
		ErrorMsg: "Sport cannot be empty",
		Label:    "Select the sport of the team",
		}, teams)
		
	away_team_id, err := controllers.GetIdByName(away_team)

	if err != nil {
		panic("error getting id")
	}

	match_date := utils.PromptGetInput(utils.PromptContent{
		ErrorMsg: "Date cannot be empty",
		Label:    "Enter the date of the match (yyyy-mm-dd)",
	})



	controllers.CreateMatch(home_team_id, away_team_id, match_date)
}