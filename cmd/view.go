/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"capi/data/store"
	"fmt"

	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: viewTable,
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func viewTable(cmd *cobra.Command, args []string) {
	if(len(args) >= 2){
		_, err := store.GetAll(args[0]+".db",args[1])
		if(err != nil){
			fmt.Printf("error getting table data: %s", err)
		}
	} else {
		store.ViewTables(args[0]+".db")
	}
}
