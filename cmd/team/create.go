/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package team

import (
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/controllers"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		createNewTeam()
	},
}

func init() {
	TeamCmd.AddCommand(createCmd)
}

func createNewTeam() {
	name := utils.PromptGetInput(utils.PromptContent{
		ErrorMsg: "Name cannot be empty",
		Label:    "Enter the name of the team",
	})

	items := controllers.GetSportNames()
	sport := utils.PromptGetSelect(utils.PromptContent{
			ErrorMsg: "Sport cannot be empty",
			Label:    "Select the sport of the team",
	}, items)

	controllers.CreateTeam(name, sport)
}