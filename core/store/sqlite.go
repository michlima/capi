package store

import (
	"database/sql"

	data "capi/core"

	_ "github.com/mattn/go-sqlite3"
)



func Open(path string, table string, keys string) (*sql.DB, error) {
	db, err := data.OpenDatabase(path)
	if err != nil {return nil,err}
	_, err = db.Exec(`PRAGMA journal_mode=WAL;`)
	if err != nil {
		return nil, err
	}

	err = data.CreateTable(db,table, keys)
	if err != nil{
		return nil, err
	}

	return db,nil
}

func ViewTables(path string) ([]string,error) {
	db, err := data.OpenDatabase(path)
	tables := []string{}
	if err != nil {
		return nil ,err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table';`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil,err
		}
		tables = append(tables, name)
	}
	return tables, nil
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

func GetAll(path string,table string) error{
	db, err := data.OpenDatabase(path)
	if err != nil {return err}
	defer db.Close()

	rows, err := db.Query(data.ViewTable(table))
	if err != nil{
		return err
	}
	defer rows.Close()
	
	cols, _ := rows.Columns()
	data.PrintHearders(cols)
	err = data.PrintRows(rows, cols)
	if(err != nil){return err}
	
	return nil
}

func GetCols(path string, table string) ([]string, error) {
	db, err := data.OpenDatabase(path)
	if err != nil {return nil,err}
	defer db.Close()

	rows, err := db.Query(data.ViewTable(table))
	if err != nil{
		return nil,err
	}
	defer rows.Close()
	cols, _ := rows.Columns()
	
	return cols,nil
}

func Get(path string, table string,filter string) error {
	db, err := data.OpenDatabase(path)
	if err != nil {return err}
	defer db.Close()
	err = data.ViewFilter(db,table,filter)
	return err
}

func Delete(path string, table string, filter string) error {
	db, err := data.OpenDatabase(path)
	if err != nil {
		return err
	}
	defer db.Close()

	err = data.Delete(db, table, filter)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTable(path string, table string) error {
	db, err := data.OpenDatabase(path)
	if err != nil {
		return err
	}
	defer db.Close()

	err = data.DropTable(db, table)
	if err != nil {
		return err
	}

	return nil
}
