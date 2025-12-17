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
	Short: "Insert data into table. Arguments: <database> <table> <cols> <values>",
	Long: `	Data needs to be inserts by argument in <database> <table> <cols> <values> For example:
	example: database myTable key,value id,someValue`,
	Run: insert,
}

func init() {
	rootCmd.AddCommand(insertCmd)
}

func insert (cmd *cobra.Command, args []string) {
	if(len(args) < 4){
		fmt.Println("4 arguments are needed. Arguments: <database> <table> <cols> <values> ")
	} else {
		err := store.Set(args[0]+".db",args[1],args[2],args[3])
		if err != nil{
			fmt.Printf("\n\n ERROR:\n'%s'\n\n",err)
		}
	}
}


