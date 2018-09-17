package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	names := strings.Split(`PID
ABN
Ent_Typ_Cd
Org_Nm
Prev_Org_Nm
Org_Nm_Change_Date
Nm_Titl_Cd
Prsn_Gvn_Nm
Prsn_Othr_Gvn_Nm
Prsn_Fmly_Nm
Nm_Sufx_Cd
Prev_Nm_Titl_Cd
Prev_Prsn_Gvn_Nm
Prev_Prsn_Othr_Gvn_Nm
Prev_Prsn_Fmly_Nm
Prev_Nm_Sufx_Cd
Nm_Change_Date
ABN_Regn_Dt
ABN_Cancn_Dt
Mn_Trdg_Nm
Prev_Mn_Trdg_Nm
Mn_Trdg_Change_Date
SON_Addr_Ln_1
SON_Addr_Ln_2
SON_Sbrb
SON_Stt
SON_Pc
SON_Cntry_Cd
SON_DPID
Prev_SON_Addr_Ln_1
Prev_SON_Addr_Ln_2
Prev_SON_Sbrb
Prev_SON_Stt
Prev_SON_Pc
Prev_SON_Cntry_Cd
Prev_SON_DPID
SON_Change_Date
Mn_Bus_Addr_Ln_1
Mn_Bus_Addr_Ln_2
Mn_Bus_Sbrb
Mn_Bus_Stt
Mn_Bus_Pc
Mn_Bus_Cntry_Cd
Mn_Bus_DPID
Prev_Mn_Bus_Addr_Ln_1
Prev_Mn_Bus_Addr_Ln_2
Prev_Mn_Bus_Sbrb
Prev_Mn_Bus_Stt
Prev_Mn_Bus_Pc
Prev_Mn_Bus_Cntry_Cd
Prev_Mn_Bus_DPID
Mn_Bus_Change_date
Ent_Eml
Prev_Ent_Eml
Ent_Eml_Change_date
Prty_Id_Blnk
GST_Regn_Dt
GST_Cancn_Dt
Mn_Indy_Clsn
Mn_Indy_Clsn_Descn
Prev_Mn_Indy_Clsn
Prev_Mn_Indy_Clsn_Descn
Mn_Indy_Change_Date
ACN
Sprsn_Ind`, "\n")

	one, _ := os.Open(os.Args[1])

	scanner := bufio.NewScanner(one)
	for scanner.Scan() {
		record := strings.SplitN(scanner.Text(), "\t", len(names))
		if len(record) < 34 {
			log.Print("Skipping short row ", scanner.Text())
			os.Exit(1)
			continue
		}

		for i := 0; i < len(names); i++ {
			fmt.Printf("%s: %s\n", names[i], record[i])
		}
		os.Exit(1)

	}

}
