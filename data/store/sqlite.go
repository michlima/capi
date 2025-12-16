package store

import (
	"database/sql"
	"fmt"

	"capi/data"

	_ "github.com/mattn/go-sqlite3"
)

func Open(path string, table string, keys string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Enable WAL mode (VERY important)
	_, err = db.Exec(`PRAGMA journal_mode=WAL;`)
	if err != nil {
		return nil, err
	}
	err = data.CreateTable(db,table, keys)
	if err != nil{
		return nil, err
	}
	return nil,nil
}

func ViewTables(path string) ( *bool,error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table';`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println("Tables in database:")
	
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		fmt.Println("-", name)
	}
	return nil , nil
}

func Set(path string, table string, keys string,values string) error {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	defer db.Close()
	err = data.Insert(db,table,keys,values)
	if(err != nil){
		return err
	}
	return nil
}

func GetAll(path string,table string) ( string, error){
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return "", err
	}
	defer db.Close()

	rows, err := db.Query(data.ViewTable(table))
	if err != nil{
		return "", err
	}
	defer rows.Close()
	
	cols, _ := rows.Columns()

	values := make([]any,len(cols))
	ptrs := make([]any, len(cols))
	for i := range values{
		ptrs[i] = &values[i]
		fmt.Printf("%s\t\t|",cols[i])
	}
	fmt.Println()
	for rows.Next() {
		rows.Scan(ptrs...)
		for _, v := range values{
			fmt.Printf("%v \t\t|", v)
		}	
		fmt.Println("")
	}
	
	return "", nil
}

func Get(db *sql.DB, key string) (string, error) {
	var value string
	err := db.QueryRow(`SELECT value FROM kv WHERE key = ?`, key).Scan(&value)
	return value, err
}
