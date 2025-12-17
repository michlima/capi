package store

import (
	"database/sql"
	"fmt"

	"capi/data"

	_ "github.com/mattn/go-sqlite3"
)



func Open(path string, table string, keys string) (*sql.DB, error) {
	db, err := data.OpenDatabase(path)
	if err != nil {return nil,err}
	defer db.Close()
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
	db, err := data.OpenDatabase(path)
	if err != nil {return nil,err}
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
	db, err := data.OpenDatabase(path)
	if err != nil {return err}
	defer db.Close()

	err = data.Insert(db,table,keys,values)
	if(err != nil){
		return err
	}
	return nil
}

func GetAll(path string,table string) ( string, error){
	db, err := data.OpenDatabase(path)
	if err != nil {return "",err}
	defer db.Close()

	rows, err := db.Query(data.ViewTable(table))
	if err != nil{
		return "", err
	}
	defer rows.Close()
	
	cols, _ := rows.Columns()
	data.PrintHearders(cols)
	data.PrintRows(rows, cols)
	
	
	return "", nil
}

func Get(path string, table string,filter string) (*string, error) {
	db, err := data.OpenDatabase(path)
	if err != nil {return nil,err}
	defer db.Close()
	err = data.ViewFilter(db,table,filter)
	return nil, err
}
