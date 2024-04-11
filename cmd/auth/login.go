/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package auth

import (
	"fmt"

	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/controllers"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login command",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		email, err := cmd.Flags().GetString("email")
		if err != nil {
			panic(err)
		}
		password, err := cmd.Flags().GetString("password")
		
		if err != nil {
			panic(err)
		}

		fmt.Println(controllers.Login(email, password))
		
		
	},
}

func init() {
	loginCmd.Flags().StringP("email", "e", "", "Email")
	loginCmd.Flags().StringP("password", "p", "", "Password")

	if err := loginCmd.MarkFlagRequired("email"); err != nil {
		fmt.Println(err)
	}

	if err := loginCmd.MarkFlagRequired("password"); err != nil {
		fmt.Println(err)
	}

	AuthCmd.AddCommand(loginCmd)

}
