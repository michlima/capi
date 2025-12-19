package interact

import (
	data "capi/core"
	"capi/core/store"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)


func UseDB (dbName string) error{
	var choice string
	for{
		tables, err := store.ViewTables(dbName+".db")
		if (err!=nil){return err}
		for i:=range tables{
			fmt.Println("- ", tables[i])
		}
		tables = append(tables, "Back")
		prompt := &survey.Select{
			Message: "Chose a table",
			Options: tables,
			Default: tables[0],
		}
		err = survey.AskOne(prompt,&choice)
		if(err !=nil){fmt.Println(err)}
		if choice == "Back" {
			return nil
		}
		err = TableInteraction(dbName,choice)
		if(err !=nil) {return err}

	}

}

func TableInteraction (dbName string, table string) error{
	var choice string
	for  {
		options := []string{
			"View",
			"Edit",
			"Add",
			"Delete",
			"Back",
		}
		prompt := &survey.Select{
			Message: "Chose a table",
			Options: options,
			Default: options[0],
		}
	
		survey.AskOne(prompt, &choice)
	
		switch choice {
		case "View":
			err := store.GetAll(dbName+".db", table)
			if(err!=nil) {return err}
		case "Add":
			availableCols,err := store.GetCols(dbName+".db",table)
			fmt.Println()
			if(err != nil) {return err}
			fmt.Println("there")
			data.PrintHearders(availableCols)
			keys := data.GetInput("> Cols : ")
			
			if(keys == "a"){
				data.PrintHearders(availableCols)
				keys=""
				for i,v:= range availableCols{
					if(i == 0){
						keys = v
					} else {
						keys = keys + ","+v
					}	
				}
			} else {
				cols:= strings.Split(keys, ",")
				data.PrintHearders(cols)
			}

			fmt.Println("EDITING:")
			
			err = addRows(dbName,table,keys)
			if(err != nil){ return err}

		case "Back":
			return nil
		}
	}	
}

func addRows (dbName string,table string,keys string) error{
	db,err := data.OpenDatabase(dbName+".db")
	if (err != nil) {return err}
	for {
		rows := data.GetInput(">")
		if(rows == "exit" ||rows == "q"){ return nil}
		err := data.Insert(db,table,keys,rows)
		if(err !=nil) {fmt.Println(err)}
		
	}
}

// func AddRows()

