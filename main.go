/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/cmd"
	"github.com/matheusnb99/sdv-m2-renf-backend-go-cobra/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}



func main() {
	cmd.Execute()
}
