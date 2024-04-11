/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/

package auth

import (
	"fmt"

	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/controllers"
	"github.com/spf13/cobra"
)

var signupCmd = &cobra.Command{
	Use:   "signup",
	Short: "Create a new user account",
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

		
		fmt.Println(controllers.SignUp(email, password))
	},
}

func init() {
	signupCmd.Flags().StringP("email", "e", "", "Email")
	signupCmd.Flags().StringP("password", "p", "", "Password")

	if err := signupCmd.MarkFlagRequired("email"); err != nil {
		fmt.Println(err)
	}

	if err := signupCmd.MarkFlagRequired("password"); err != nil {
		fmt.Println(err)
	}


	AuthCmd.AddCommand(signupCmd)

}
