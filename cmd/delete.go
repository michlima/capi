/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"capi/data/store"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	tableToDelete  string
	filterToDelete string
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <database>",
	Short: "Delete tables or rows from a database",
	Long: `Delete tables or rows from a database.

To delete a table:
  capi delete <database> -t <tableName>

To delete rows from a table:
  capi delete <database> -t <tableName> -f "column=value"`,
	Run: delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&tableToDelete, "table", "t", "", "Table to delete from")
	deleteCmd.Flags().StringVarP(&filterToDelete, "filter", "f", "", "Filter to select rows for deletion (e.g., 'id=1')")
}

func delete(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: database name is required")
		return
	}
	dbName := args[0] + ".db"

	if tableToDelete == "" {
		fmt.Println("Error: table name is required. Use -t flag.")
		return
	}

	if filterToDelete != "" {
		// Delete rows
		if err := store.Delete(dbName, tableToDelete, filterToDelete); err != nil {
			fmt.Printf("Error deleting rows: %v\n", err)
		} else {
			fmt.Printf("Successfully deleted rows from table '%s'\n", tableToDelete)
		}
	} else {
		// Delete table
		if err := store.DeleteTable(dbName, tableToDelete); err != nil {
			fmt.Printf("Error deleting table: %v\n", err)
		} else {
			fmt.Printf("Successfully deleted table '%s'\n", tableToDelete)
		}
	}
}
