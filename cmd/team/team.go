/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package team

import (
	"github.com/spf13/cobra"
)

// teamCmd represents the team command
var TeamCmd = &cobra.Command{
	Use:   "team",
	Short: "A brief description of your command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
}
