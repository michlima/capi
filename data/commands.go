package data

import (
	"database/sql"
	"fmt"
	"strings"
)

func CreateTable(db *sql.DB,table string, keys string) (error) {
	cols := strings.Split(keys, ",")

    if len(cols) == 0 {
        return fmt.Errorf("no columns provided")
    }

    // First column is primary key
    columnDefs := []string{fmt.Sprintf("%s TEXT PRIMARY KEY", cols[0])}

    // Remaining columns
    for _, col := range cols[1:] {
        columnDefs = append(columnDefs, fmt.Sprintf("%s TEXT", col))
    }

    cmd := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);",
        table,
        strings.Join(columnDefs, ", "),
    )

    _, err := db.Exec(cmd)
    if err != nil {
        fmt.Println("returning error:", err)
        return err
    }

    return nil
}

func Insert(db *sql.DB,table string, keys string,values string ) (error) {	
	var parts = strings.Split(values, ",")
	placeholders := make([]string, len(parts))
	for i := range parts {
		placeholders[i] = fmt.Sprintf("$%d", i+1) // PostgreSQL style placeholders
	}
	placeholderStr := strings.Join(placeholders, ",")	
	
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", table, keys, placeholderStr)
	
	var vls = make([]interface{},len(parts))
	for i,v := range parts {
		vls[i] = strings.TrimSpace(v)
	}

	_, err := db.Exec(query, vls...)
	return err
}


func ViewTable(tableName string) string{
	var cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	return cmd
}