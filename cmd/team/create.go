/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package team

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/controllers"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		createNewTeam()

	},
}

type promptContent struct {
	errorMsg string
	label    string
}



func init() {
	TeamCmd.AddCommand(createCmd)

}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}


func promptGetSelect(pc promptContent) string {
	items := controllers.GetSportNames()
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func createNewTeam() {
	name := promptGetInput(promptContent{
		errorMsg: "Name cannot be empty",
		label:    "Enter the name of the team",
	})

	sport := promptGetSelect(promptContent{
		errorMsg: "Sport cannot be empty",
		label:    "Select the sport of the team",
	})

	controllers.CreateTeam(name, sport)
}