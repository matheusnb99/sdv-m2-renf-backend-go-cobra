/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/cmd/auth"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/cmd/match"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/cmd/team"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/cmd/tournament_"

	"github.com/spf13/cobra"
)


func addSubCommandPalettes() {
	RootCmd.AddCommand(auth.AuthCmd)
	RootCmd.AddCommand(tournament_.TournamentCmd)
	RootCmd.AddCommand(match.MatchCmd)
	RootCmd.AddCommand(team.TeamCmd)


}


// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sdv-m2-renf-backend-go-cobra",
	Short: "This is a backend cli for the TourNament project, using Go and Cobra.",
	Long: `This is a backend cli for the TourNament project, using Go and Cobra. 
	You can use this cli to manage the database, create users, authenticate users and more.

	It uses GORM as the ORM and JWT and Bcrypt for authentication.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommandPalettes()
}


