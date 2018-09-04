package main

import (
	"github.com/SmallBusinessVictoria-ABR-ETL/tools"
	"log"
	"os"
)

func main() {
	one, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	two, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	newRecords, err := os.Create(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	tools.Diff(one, two, newRecords, os.Args[4])
}
