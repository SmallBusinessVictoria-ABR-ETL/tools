package main

import (
	"github.com/SmallBusinessVictoria-ABR-ETL/tools"
	"os"
)

func main() {

	f, _ := os.Open(os.Args[1])
	tools.Validate(f)

}
