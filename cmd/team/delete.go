/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package team

import (
	"fmt"

	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/controllers"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a team by name",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			panic("error getting name")
		}

		controllers.DeleteTeam(name);
	},
}




func init() {
	deleteCmd.Flags().StringP("name", "n", "", "Name")

	if err := deleteCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}

	TeamCmd.AddCommand(deleteCmd)
}
