/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"capi/data/store"
	"fmt"

	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert data into table",
	Long: `Data needs to be inserts by argument in <database> <table> <cols> <values> For example:
	example: database myTable key,value id,someValue`,
	Run: func(cmd *cobra.Command, args []string) {
		err := store.Set(args[0]+".db",args[1],args[2],args[3])
		if err != nil{
			fmt.Printf("\n\n ERROR:\n'%s'\n\n",err)
		}
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	
	

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
