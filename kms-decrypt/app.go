package main

import (
	"fmt"
	"github.com/SmallBusinessVictoria-ABR-ETL/tools"
	"os"
)

func main() {
	fmt.Print(tools.Decrypt(os.Args[1]))
}
