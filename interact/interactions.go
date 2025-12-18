package interact

import (
	"capi/core/store"
	"fmt"

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

		case "Back":
			return nil
		}
	}	
}

// func AddRows()

