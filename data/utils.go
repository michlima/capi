package data

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase(path string) (*sql.DB, error){
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		fmt.Println("Error opening database: %s",err)
		return nil, err
	}
	return db, nil
}

func PrintHearders(cols[]string ){
	fmt.Printf("\n ")
	for range cols{
		for i := 0; i < 16	; i++{
			fmt.Printf("-")
		}
	}
	fmt.Println()
	for i := range cols{
		if(i == 0) {fmt.Print("|")}
		fmt.Printf("%-15s|",cols[i])
	}
	fmt.Printf("\n ")
	for range cols{
		for i := 0; i < 16	; i++{
			fmt.Printf("-")
		}
	}
	fmt.Println()
}

func PrintRows(rows *sql.Rows, cols []string) error{
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