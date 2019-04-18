package main

import (
	"bytes"
	"fmt"
	"github.com/SmallBusinessVictoria-ABR-ETL/tools"
	"log"
	"os"
	"text/template"
)

func main() {

	data := struct {
		Args []string
		Arg1 string
		Arg2 string
	}{
		Args: os.Args[2:],
		Arg1: os.Args[2],
		Arg2: os.Args[3],
	}

	var sql bytes.Buffer
	tpl, err := template.ParseFiles(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(&sql, data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(sql.Bytes()))
	tools.Query(sql.String(), "s3://sbv-abr-etl/custom-extract/"+data.Arg1+"-"+data.Arg2)

	//fmt.Print("\n-----\n\ns3://sbv-abr-etl/custom-extract/"+data.Arg1+".csv\n\n-----\n")
}
