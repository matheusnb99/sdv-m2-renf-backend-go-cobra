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
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
