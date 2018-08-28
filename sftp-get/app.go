package main

import (
	"github.com/SmallBusinessVictoria-ABR-ETL/tools"
	"os"
)

func main() {
	tools.SFTPGet(os.Args[1], os.Args[2])
}
