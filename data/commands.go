package data

import (
	"database/sql"
	"fmt"
	"strings"
)

func CreateTable(name string)(string) {
	cmd := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (key TEXT PRIMARY KEY, value TEXT);", name)
	return cmd
}

func Insert(db *sql.DB,table string, keys string,values string ) (error) {	
	var parts = strings.Split(values, ",")
	placeholders := make([]string, len(parts))
	for i := range parts {
		placeholders[i] = fmt.Sprintf("$%d", i+1) // PostgreSQL style placeholders
	}
	placeholderStr := strings.Join(placeholders, ", ")	
	
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