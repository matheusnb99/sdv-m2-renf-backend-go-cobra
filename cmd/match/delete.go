/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package match

import (
	"fmt"
	"strconv"

	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/controllers"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a match by id",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			panic("error getting id")
		}

		// Convert id from string to uint
		matchID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			panic("error converting id to uint")
		}

		controllers.DeleteMatch(uint(matchID))
	},
}

func init() {
	deleteCmd.Flags().StringP("id", "i", "", "Id")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}

	MatchCmd.AddCommand(deleteCmd)

}
