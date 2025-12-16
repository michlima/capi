/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"capi/data/store"
	"fmt"

	"github.com/spf13/cobra"
)

var t string
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
}

func viewTable(cmd *cobra.Command, args []string) {
	if(len(t) > 1){
		_, err := store.GetAll(args[0]+".db",t)
		if(err != nil){
			fmt.Printf("error getting table data: %s", err)
		}
	} else {
		store.ViewTables(args[0]+".db")
	}
}
