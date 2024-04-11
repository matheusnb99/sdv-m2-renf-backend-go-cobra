/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package team

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
	Short: "List the teams",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		
		lineFlag, _ := cmd.Flags().GetBool("line")
		interactiveFlag, _ := cmd.Flags().GetBool("interactive")
			
		teamNames := controllers.GetTeamNames()

		switch {
		case interactiveFlag:
			selectedTeam := utils.PromptGet(teamNames)
			fmt.Println("Selected team:", selectedTeam)
			matches := controllers.GetMatchesByName(selectedTeam)

			fmt.Println("id | scoreTeamAway | scoreTeamLocal | win | matchDate")
			fmt.Println(matches)
		case lineFlag:
			fmt.Println(strings.Join(teamNames, "\n"))
		default:
			fmt.Println(strings.Join(teamNames, ", "))
		}
	},
}

func init() {
	listCmd.Flags().BoolP("line", "l", false, "Print list as a line")
	listCmd.Flags().BoolP("interactive", "i", false, "Print interactive list")
	TeamCmd.AddCommand(listCmd)
}



