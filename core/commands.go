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

func Insert(db *sql.DB,table string, cols string,values string ) (error) {	
	var parts = strings.Split(values, ",")
	placeholders := make([]string, len(parts))
	for i := range parts {
		placeholders[i] = fmt.Sprintf("$%d", i+1) // PostgreSQL style placeholders
	}
	placeholderStr := strings.Join(placeholders, ",")	
	
	query := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", table, cols, placeholderStr)
	
	var vls = make([]interface{},len(parts))
	for i,v := range parts {
		vls[i] = strings.TrimSpace(v)
	}

	_, err := db.Exec(query, vls...)
	return err
}

func ViewFilter(db *sql.DB, table string, filter string) error {
    parts := strings.Split(filter, ",")
    q := fmt.Sprintf("SELECT * FROM %s WHERE ", table)
    var vls []interface{}

    for i, p := range parts {
        kv := strings.SplitN(p, "=", 2)
        if i > 0 {
            q += " OR "
        }
        q += fmt.Sprintf("[%s]=?", kv[0])
        vls = append(vls, strings.TrimSpace(kv[1]))
    }

    rows, err := db.Query(q, vls...)
    if err != nil {
        return err
    }
    defer rows.Close()

    cols, err := rows.Columns()
    if err != nil {
        return err
    }

    PrintHearders(cols)
	err = PrintRows(rows, cols)
	if(err != nil){return err}

    return nil
}


func ViewTable(tableName string) string{
	var cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	return cmd
}

func Delete(db *sql.DB, table string, filter string) error {


	parts := strings.Split(filter, ",")


	q := fmt.Sprintf("DELETE FROM %s WHERE ", table)


	var vls []interface{}





	for i, p := range parts {


		kv := strings.SplitN(p, "=", 2)


		if i > 0 {


			q += " OR "


		}


		q += fmt.Sprintf("[%s]=?", kv[0])


		vls = append(vls, strings.TrimSpace(kv[1]))


	}





	_, err := db.Exec(q, vls...)


	return err


}

func DropTable(db *sql.DB, table string) error {


	cmd := fmt.Sprintf("DROP TABLE IF EXISTS %s", table)


	_, err := db.Exec(cmd)


	return err


}

