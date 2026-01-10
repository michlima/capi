package uploads

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func uploadCsv(db string, table string, path string) {
	 file, err := os.Open("year.csv")
        if err != nil {
                panic(err)
        }
        defer file.Close()
        reader := csv.NewReader(file)
        var today any
        for {
                record, err := reader.Read()
                if err == io.EOF {
                        break
                }
                if err != nil {
                        panic(err)
                }
                var complete = record[6]
                if(complete == "0") {
                        today = record
                        break
                }
        }
        fmt.Println(today)
}