/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/initializers"

	"github.com/spf13/cobra"
)



var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database",
	Long: `This command initializes the database, creating the tables and columns.`,
	Run: func(cmd *cobra.Command, args []string) {
			initializers.LoadEnvVariables()
			fmt.Println("Connecting to database...")
			initializers.ConnectToDb()
			fmt.Println("Connected to database")
			fmt.Println("Syncing")
			initializers.SyncDb()
			fmt.Println("Database initialized")
		
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
