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

var table string
var cols string

var ctCmd = &cobra.Command{
	Use:   "ct",
	Short: "Create table in database (first argument is the name of the database)",
	Long: `Create table in database (first argument is the name of the database). For example:

	example: capi ct <databaseName> -t <tableName> -c <columnsOnTable>
	applied: capi ct myApp -t users -c id,name,lastname,phonesNumber,streetAdr
	`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := store.Open(args[0]+".db", table, cols)
		if err != nil{
			fmt.Println("printing from main function")
			fmt.Print(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ctCmd)
	
	ctCmd.Flags().StringVarP(&table,"table","t","defaultTable","name table you want to create")
	ctCmd.Flags().StringVarP(&cols,"cols","c","key,value","columns on table (first column is primary key)")
	
}
