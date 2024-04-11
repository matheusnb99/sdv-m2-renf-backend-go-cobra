/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package tournament_

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	TournamentCmd.AddCommand(createCmd)

}
