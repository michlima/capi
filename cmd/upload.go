/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"capi/core/store"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)
var typ string
var path string
var tab string
var dab string
func arrToString(arr []string) string {
	var str = ""
	for i,data := range arr{
		if(i == 0){
			str = str + data
		} else {
			str =  str+","+ data
		}
	}
	return str 
}
// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload data from a file into a database table",
	Long: `Upload data from a specified file into a database table.

This command currently supports CSV files. It reads the contents of the file,
retrieves the column names from the target table, and inserts each row into the
database.

Examples:

  # Upload a CSV file into a specific table
  capi upload --type csv --path /path/to/file.csv --db mydatabase --tb mytable

Flags:

  --type, -y   Type of the input file (e.g., csv)
  --path, -p   Path to the input file
  --db         Database to use
  --tb         Target table to insert data

Note:

- The file must be in CSV format if --type is set to "csv".
- The target table must already exist in the database.
- All required flags (--type, --path, --tb) must be provided.`,
	Run: func(cmd *cobra.Command, args []string) {
		if(typ == "csv"){
			uploadCsv(dab, tab, path)
		}
	},
}

func uploadCsv(db string, table string, path string) error {
	 file, err := os.Open(path)
	 	fmt.Println("1")
        if err != nil {
			fmt.Println("Error opening file")
            panic(err)
        }
        defer file.Close()
        reader := csv.NewReader(file)
		cols,err := store.GetCols(db+".db",table)
		if(err != nil){
			fmt.Println(err)
			return err
		}
		
        for {
			// get values
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			// upload to sqlite
			var values = arrToString(record)
			var colStr = arrToString(cols)
			// fmt.Println(colStr)
			fmt.Println(values)
			store.Set(db+".db",table,colStr,values)
        }
		return nil
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&typ, "type", "y", "", "type of file")
    uploadCmd.Flags().StringVarP(&path, "path", "p", "", "path to file")
    uploadCmd.Flags().StringVar(&dab, "db", "", "database being used")
    uploadCmd.Flags().StringVar(&tab, "tb", "", "inserting to this table")

    uploadCmd.MarkFlagRequired("path")
    uploadCmd.MarkFlagRequired("tb")
    uploadCmd.MarkFlagRequired("type")
}
