/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package auth

import (
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Commands related to authentication",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
