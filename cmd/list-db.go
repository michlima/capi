/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"capi/core"
	"fmt"

	"github.com/spf13/cobra"
)

// listDbCmd represents the list-db command
var listDbCmd = &cobra.Command{
	Use:   "list-db",
	Short: "List all available SQLite databases in the storage directory",
	Long:  `This command scans the storage directory and lists all files ending with '.db', which are treated as SQLite database files.`,
	Run:   listDB,
}

func init() {
	rootCmd.AddCommand(listDbCmd)
}

func listDB(cmd *cobra.Command, args []string) {
	databases, err := data.ListDatabases()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(databases) == 0 {
		fmt.Println("No SQLite databases found in the storage directory.")
		return
	}

	fmt.Println("Available SQLite Databases:")
	for _, dbName := range databases {
		fmt.Printf("- %s\n", dbName)
	}
}
