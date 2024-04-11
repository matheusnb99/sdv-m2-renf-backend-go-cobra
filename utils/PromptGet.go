package utils

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func PromptGet(items []string) string {

	prompt := promptui.Select{
		Label: "Select Day",
		Items: items,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}
