/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"capi/data/store"
	"fmt"

	"github.com/spf13/cobra"
)

// ctCmd represents the ct command
var ctCmd = &cobra.Command{
	Use:   "ct",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := store.Open(args[0]+".db", args[1])
		if err != nil{
			fmt.Println("printing from main function")
			fmt.Print(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ctCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ctCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ctCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
