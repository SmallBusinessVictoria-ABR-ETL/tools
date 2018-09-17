package tools

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var namesCombined = []string{
	"PID",                     // 0
	"ABN",                     // 1
	"Ent_Typ_Cd",              // 2
	"Org_Nm",                  // 3
	"Prev_Org_Nm",             // 4
	"Org_Nm_Change_Date",      // 5
	"Nm_Titl_Cd",              // 6
	"Prsn_Gvn_Nm",             // 7
	"Prsn_Othr_Gvn_Nm",        // 8
	"Prsn_Fmly_Nm",            // 9
	"Nm_Sufx_Cd",              // 10
	"Prev_Nm_Titl_Cd",         // 11
	"Prev_Prsn_Gvn_Nm",        // 12
	"Prev_Prsn_Othr_Gvn_Nm",   // 13
	"Prev_Prsn_Fmly_Nm",       // 14
	"Prev_Nm_Sufx_Cd",         // 15
	"Nm_Change_Date",          // 16
	"ABN_Regn_Dt",             // 17
	"ABN_Cancn_Dt",            // 18
	"Mn_Trdg_Nm",              // 19
	"Prev_Mn_Trdg_Nm",         // 20
	"Mn_Trdg_Change_Date",     // 21
	"SON_Addr_Ln_1",           // 22
	"SON_Addr_Ln_2",           // 23
	"SON_Sbrb",                // 24
	"SON_Stt",                 // 25
	"SON_Pc",                  // 26
	"SON_Cntry_Cd",            // 27
	"SON_DPID",                // 28
	"Prev_SON_Addr_Ln_1",      // 29
	"Prev_SON_Addr_Ln_2",      // 30
	"Prev_SON_Sbrb",           // 31
	"Prev_SON_Stt",            // 32
	"Prev_SON_Pc",             // 33
	"Prev_SON_Cntry_Cd",       // 34
	"Prev_SON_DPID",           // 35
	"SON_Change_Date",         // 36
	"Mn_Bus_Addr_Ln_1",        // 37
	"Mn_Bus_Addr_Ln_2",        // 38
	"Mn_Bus_Sbrb",             // 39
	"Mn_Bus_Stt",              // 40
	"Mn_Bus_Pc",               // 41
	"Mn_Bus_Cntry_Cd",         // 42
	"Mn_Bus_DPID",             // 43
	"Prev_Mn_Bus_Addr_Ln_1",   // 44
	"Prev_Mn_Bus_Addr_Ln_2",   // 45
	"Prev_Mn_Bus_Sbrb",        // 46
	"Prev_Mn_Bus_Stt",         // 47
	"Prev_Mn_Bus_Pc",          // 48
	"Prev_Mn_Bus_Cntry_Cd",    // 49
	"Prev_Mn_Bus_DPID",        // 50
	"Mn_Bus_Change_date",      // 51
	"Ent_Eml",                 // 52
	"Prev_Ent_Eml",            // 53
	"Ent_Eml_Change_date",     // 54
	"Prty_Id_Blnk",            // 55
	"GST_Regn_Dt",             // 56
	"GST_Cancn_Dt",            // 57
	"Mn_Indy_Clsn",            // 58
	"Mn_Indy_Clsn_Descn",      // 59
	"Prev_Mn_Indy_Clsn",       // 60
	"Prev_Mn_Indy_Clsn_Descn", // 61
	"Mn_Indy_Change_Date",     // 62
	"ACN",       // 63
	"Sprsn_Ind", // 64
}

var namesRaw = []string{
	"PID",                // 0
	"ABN",                // 1
	"Ent_Typ_Cd",         // 2
	"Org_Nm",             // 3
	"Nm_Titl_Cd",         // 4
	"Prsn_Gvn_Nm",        // 5
	"Prsn_Othr_Gvn_Nm",   // 6
	"Prsn_Fmly_Nm",       // 7
	"Nm_Sufx_Cd",         // 8
	"ABN_Regn_Dt",        // 9
	"ABN_Cancn_Dt",       // 10
	"Mn_Trdg_Nm",         // 11
	"SON_Addr_Ln_1",      // 12
	"SON_Addr_Ln_2",      // 13
	"SON_Sbrb",           // 14
	"SON_Stt",            // 15
	"SON_Pc",             // 16
	"SON_Cntry_Cd",       // 17
	"SON_DPID",           // 18
	"Mn_Bus_Addr_Ln_1",   // 19
	"Mn_Bus_Addr_Ln_2",   // 20
	"Mn_Bus_Sbrb",        // 21
	"Mn_Bus_Stt",         // 22
	"Mn_Bus_Pc",          // 23
	"Mn_Bus_Cntry_Cd",    // 24
	"Mn_Bus_DPID",        // 25
	"Ent_Eml",            // 26
	"Prty_Id_Blnk",       // 27
	"GST_Regn_Dt",        // 28
	"GST_Cancn_Dt",       // 29
	"Mn_Indy_Clsn",       // 30
	"Mn_Indy_Clsn_Descn", // 31
	"ACN",                // 32
	"Sprsn_Ind",          // 33
}

func View(one io.Reader) {
	scanner := bufio.NewScanner(one)
	lnames := len(namesCombined)
	for scanner.Scan() {
		record := strings.SplitN(scanner.Text(), "\t", lnames)
		if len(record) < lnames {
			log.Print("Skipping short row ", scanner.Text())
			ViewRow(record)
			os.Exit(1)
			continue
		}

		ViewRow(record)
		os.Exit(1)
	}

}
func ViewPid(one io.Reader, pid string) {
	scanner := bufio.NewScanner(one)
	lnames := len(namesCombined)
	lnamesRaw := len(namesRaw)
	for scanner.Scan() {
		record := strings.SplitN(scanner.Text(), "\t", lnames)
		if len(record) >= 1 {
			if record[1] == pid {
				if len(record) == lnames {
					ViewRow(record)
					os.Exit(1)
					continue
				}
				if len(record) == lnamesRaw {
					ViewRowRaw(record)
					os.Exit(1)
					continue
				}

			}
		}

		ViewRow(record)
		os.Exit(1)
	}

}
func Validate(one io.Reader) {
	scanner := bufio.NewScanner(one)
	lnames := len(namesCombined)
	for scanner.Scan() {
		record := strings.SplitN(scanner.Text(), "\t", lnames)
		if len(record) != lnames {
			log.Print("Found short row\n")
			ViewRow(record)
			os.Exit(1)
			continue
		}
	}
}

func ViewRow(row []string) {
	for i := 0; i < minLen(namesCombined, row); i++ {
		fmt.Printf("%s: %s\n", namesCombined[i], row[i])
	}
	if len(namesCombined) > len(row) {
		fmt.Print("...\n")
	}
}
func ViewRowRaw(row []string) {
	for i := 0; i < minLen(namesRaw, row); i++ {
		fmt.Printf("%s: %s\n", namesRaw[i], row[i])
	}
	if len(namesRaw) > len(row) {
		fmt.Print("...\n")
	}
}
func minLen(a []string, b []string) int {
	if len(a) > len(b) {
		return len(b)
	}
	return len(a)
}
