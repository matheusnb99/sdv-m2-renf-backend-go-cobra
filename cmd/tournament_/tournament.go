/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package tournament_

import (
	"github.com/spf13/cobra"
)

// tournamentCmd represents the tournament command
var TournamentCmd = &cobra.Command{
	Use:   "tournament",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
