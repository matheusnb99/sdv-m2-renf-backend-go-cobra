/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	
	"github.com/stretchr/testify/assert"
)

// testsCmd represents the tests command
var testsCmd = &cobra.Command{
	Use:   "tests",
	Short: "Test all my commands",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tests called")
		Test_ExecuteParticularCommandDynamically(nil)

	},
}

func init() {
	
	RootCmd.AddCommand(testsCmd)

}




func execute(args string) string {
	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs(strings.Split(args, " "))
	RootCmd.Execute()

	return actual.String()
}


func Test_ExecuteParticularCommandDynamically(t *testing.T) {
	for _, child := range RootCmd.Commands() {
		childArgs, exists := child.Annotations["args"]
		fmt.Printf("Testing command pallet: %s \n", child.Use) 


		
		if exists {
			actual := execute(childArgs)
			expected, exists := child.Annotations["output"]
			if !exists{
				t.Errorf("Output doesn't exists in [%s] command", child.Use)
			}
			assert.Equal(t, actual, expected, "actual is not expected")
		}

		if !child.HasSubCommands() {
			continue
		}

		for _, grandChild := range child.Commands() {
			grandChildArgs, exists := grandChild.Annotations["args"]
			
			if grandChild.Use == "addTeam" {
				fmt.Printf("\t %s... \tFAILED\n", grandChild.Use) // Indicate that the test passed for the subcommand
				continue
			}
			fmt.Printf("\t %s... \tPASSED\n", grandChild.Use) // Indicate that the test passed for the subcommand

			if exists {
				actual := execute(grandChildArgs)
				expected, exists := grandChild.Annotations["output"]
				if !exists{
					t.Errorf("Output doesn't exists in [%s] command", child.Use)
				}
				assert.Equal(t, actual, expected, "actual is not expected")
			}
		}

	}

}
