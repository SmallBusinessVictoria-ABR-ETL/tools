package main

import (
	"fmt"
	"github.com/SmallBusinessVictoria-ABR-ETL/tools"
	"log"
	"os"
)

func main() {
	plain, err := tools.Decrypt(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(plain)
}
