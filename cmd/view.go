/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	data "capi/core"
	"capi/core/store"
	"fmt"

	"github.com/spf13/cobra"
)

var t string
var f string
// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View database table and tables",
	Long: `View database table and tables. For example:
view database: 	capi view <database>
view table   :	capi view -t <table>
`,
	Run: viewTable,
}

func init() {
	rootCmd.AddCommand(viewCmd)
	viewCmd.Flags().StringVarP(&t,"table","t","","define table you want to view")
	viewCmd.Flags().StringVarP(&f,"filter","f","","filter table to see specific rows")
}

func viewTable(cmd *cobra.Command, args []string) {
	if (len(args) == 0) {
		fmt.Println("Minimum of one argument required: view <database>")
		return
	}
	if(len(f) > 1){
		if(len(t) < 1){
			fmt.Println("Table needs to be defined with flag '-t <table>'")
			return
		} else {
			err := store.Get(args[0]+".db", t, f)
			if err != nil {
				fmt.Printf("\n Error: %s\n", err)
			}
		}
	} else if(len(t) > 1){
		err := store.GetAll(args[0]+".db",t)
		if(err != nil){
			fmt.Printf("error getting table data: %s", err)
		}
	}  else {
		tables,err := store.ViewTables(args[0]+".db")
		if err != nil {
			fmt.Printf("\n Error: %s\n", err)
			return
		}
		data.PrintTables(tables)
	}
}
