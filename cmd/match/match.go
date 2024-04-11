/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package match

import (
	"github.com/spf13/cobra"
)

// matchCmd represents the match command
var MatchCmd = &cobra.Command{
	Use:   "match",
	Short: "Commands related to match",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
