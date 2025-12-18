package data

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func ListDatabases() ([]string, error) {
	files, err := ioutil.ReadDir("./storage")
	if err != nil {
		return nil, fmt.Errorf("error reading storage directory: %w", err)
	}

	var databases []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".db") {
			databases = append(databases, strings.TrimSuffix(file.Name(), ".db"))
		}
	}

	return databases, nil
}

func OpenDatabase(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./storage/"+path)
	if err != nil {
		fmt.Printf("Error opening database: %s\n", err)
		return nil, err
	}
	return db, nil
}

func PrintTables(tables []string){
	for i:=range tables{
		fmt.Println("- ", tables[i])
	}
}

func PrintHearders(cols []string) {
	fmt.Printf("\n ")
	for range cols {
		for i := 0; i < 16; i++ {
			fmt.Printf("-")
		}
	}
	fmt.Println()
	for i := range cols {
		if i == 0 {
			fmt.Print("|")
		}
		fmt.Printf("%-15s|", cols[i])
	}
	fmt.Printf("\n ")
	for range cols {
		for i := 0; i < 16; i++ {
			fmt.Printf("-")
		}
	}
	fmt.Println()
}

func PrintRows(rows *sql.Rows, cols []string) error {
	for rows.Next() {
		fmt.Print("|")
		vals := make([]interface{}, len(cols))
		ptrs := make([]interface{}, len(cols))
		for i := range vals {
			ptrs[i] = &vals[i]
		}

		if err := rows.Scan(ptrs...); err != nil {
			return err
		}

		for _, val := range vals {
			fmt.Printf("%-15v|", val)
		}
		fmt.Println()
	}
	return nil
}