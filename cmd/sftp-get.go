package main

import (
	"bitbucket.org/portable/sbv-abr-etl"
	"os"
)

func main()  {
	sbv_abr_etl.SFTPGet(os.Args[1], os.Args[2])
}