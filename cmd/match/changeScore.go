/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package match

import (
	"fmt"
	"strconv"

	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/controllers"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/utils"
	"github.com/spf13/cobra"
)

var changeScoreCmd = &cobra.Command{
	Use:   "changeScore",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		matchID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			fmt.Println("Error: Invalid ID format")
			return
		}
		changeScoreInteractive(uint(matchID))

	},
}

func init() {
	changeScoreCmd.Flags().StringP("id", "i", "", "Id")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
	MatchCmd.AddCommand(changeScoreCmd)

}


func changeScoreInteractive(id uint) {
	home_team, away_team := controllers.GetMatchTeamsNames(id)

	teams := []string{home_team, away_team}

	chosen_team := utils.PromptGetSelect(utils.PromptContent{
			ErrorMsg: "Sport cannot be empty",
			Label:    "Select the sport of the team",
	}, teams)
	chosen_team_id, err := controllers.GetIdByName(chosen_team)

	if err != nil {
		panic("error getting id")
	}


	new_score := utils.PromptGetInput(utils.PromptContent{
		ErrorMsg: "Score cannot be empty",
		Label:    "Enter new score",
	})

	num, err := strconv.ParseUint(new_score, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("changing score of team", chosen_team_id, "to", num)

	yes_no_array := []string{"yes", "no"}

	is_it_over := utils.PromptGetSelect(utils.PromptContent{
			ErrorMsg: "Anwser cannot be empty",
			Label:    "Is the match over?",
	}, yes_no_array)

	if is_it_over == "yes" {
		controllers.SetMatchOver(id)
	}


	score := uint(num)

	controllers.ChangeScoreOfTeam(id, chosen_team_id, score)


}