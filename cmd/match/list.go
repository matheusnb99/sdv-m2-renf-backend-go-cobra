/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package match

import (
	"fmt"
	"strings"

	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/controllers"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the matches",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		
		interactiveFlag, _ := cmd.Flags().GetBool("interactive")

		matchesList, err := controllers.GetMatchWithTeams()

		if err != nil {
			fmt.Println("Error getting matches")
			return
		}

		matches := []string{}
		for _, match := range matchesList {
			matches = append(matches, controllers.MatchDetailToString(match))
		}
		switch {
			case interactiveFlag:
				selectedTeam := utils.PromptGet(matches)
				fmt.Println("Selected team:", selectedTeam)
				matches := controllers.GetMatchesByName(selectedTeam)

				fmt.Println("id | scoreTeamAway | scoreTeamLocal | win | matchDate")
				fmt.Println(matches)
			default:
				fmt.Println(strings.Join(matches, "\n"))
			}
	},
}

func init() {
	listCmd.Flags().BoolP("interactive", "i", false, "Print interactive list")
	MatchCmd.AddCommand(listCmd)
}
