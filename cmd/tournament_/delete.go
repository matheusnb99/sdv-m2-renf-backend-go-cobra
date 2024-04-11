/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package tournament_

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}

func init() {
	TournamentCmd.AddCommand(deleteCmd)

	
}
