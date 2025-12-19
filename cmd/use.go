/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	data "capi/core"
	"capi/interact"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: interactive,
}

func init() {
	rootCmd.AddCommand(useCmd)
}

func interactive(cmd *cobra.Command, args []string) {
	options := []string{}
	var choice string
	dbs, err := data.ListDatabases()
	if(err!=nil) { fmt.Println("Error: ", err) }
	
	options = append(options, dbs...)
	
	
	prompt := &survey.Select{
		Message: "which database to use?",
		Options: options,
		Default: options[0],
	}
	
	err = survey.AskOne(prompt, &choice)
	if err != nil {
		fmt.Println(err)
		return 
	}
	err = interact.UseDB(choice)
	if err != nil {
		fmt.Println(err)
		return 
	}
	
}
