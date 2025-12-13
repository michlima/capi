/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to firestore database",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		getConfigs()
	},
}
func getConfigs(){
	config := make(map[string]string)
	file, err := os.OpenFile("c/c.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    // Ensure the file is closed after the function completes
    defer file.Close()
	fields := []string{"apiKey","authDomain","projectId","storageBucket","messagingSenderId","appId","measurementId"}
	for _, field := range fields {
		textInput := pterm.DefaultInteractiveTextInput
		prompt := fmt.Sprintf("Enter Firestore %s:", field)
		value, _ := textInput.Show(prompt)
		config[field] = value
	}
	
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("FIELD,SECRET\n")
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}
	for key, field := range config{
		_, err = writer.WriteString(fmt.Sprintf("%s,%s\n", key, field))
		if err != nil {
			fmt.Println("Error writing to buffer:", err)
			return
		}
		err = writer.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer:", err)
			return
		}
	}
	fmt.Printf("READY TO CONNECT!")
}	

func init() {
	rootCmd.AddCommand(connectCmd)
	
}

