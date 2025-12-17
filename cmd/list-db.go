/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// listDbCmd represents the list-db command
var listDbCmd = &cobra.Command{
	Use:   "list-db",
	Short: "List all available SQLite databases in the current directory",
	Long:  `This command scans the current directory and lists all files ending with '.db', which are treated as SQLite database files.`,
	Run: listDB,
}

func init() {
	rootCmd.AddCommand(listDbCmd)
}
func listDB(cmd *cobra.Command, args []string) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading current directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Available SQLite Databases:")
	found := false
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".db") {
			fmt.Printf("- %s\n", strings.TrimSuffix(file.Name(), ".db"))
			found = true
		}
	}

	if !found {
		fmt.Println("No SQLite databases found in the current directory.")
	}
}

