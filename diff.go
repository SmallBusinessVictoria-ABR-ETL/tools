package tools

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	OrgNameHash     []byte
	OrgName         []string
	NameHash        []byte
	Name            []string
	TradingNameHash []byte
	TradingName     []string
	SONAddressHash  []byte
	SONAddress      []string
	BusAddressHash  []byte
	BusAddress      []string
	EmailHash       []byte
	Email           []string
	IndustryHash    []byte
	Industry        []string
}

func hashRecord(r []string) Record {

	// PID	integer
	// ABN	varchar(20)
	// Ent_Typ_Cd	varchar(3)

	// Org_Nm	varchar(200)

	// Nm_Titl_Cd	varchar(12)  4
	// Prsn_Gvn_Nm	varchar(40)
	// Prsn_Othr_Gvn_Nm	varchar(100)
	// Prsn_Fmly_Nm	varchar(40)
	// Nm_Sufx_Cd	varchar(5)    8

	// ABN_Regn_Dt	varchar(8)
	// ABN_Cancn_Dt	varchar(8)

	// Mn_Trdg_Nm	varchar(200)   11

	// SON_Addr_Ln_1	varchar(38)   12
	// SON_Addr_Ln_2	varchar(38)
	// SON_Sbrb	varchar(46)
	// SON_Stt	varchar(3)
	// SON_Pc	varchar(12)
	// SON_Cntry_Cd	varchar(3)
	// SON_DPID	integer                 18

	// Mn_Bus_Addr_Ln_1	varchar(38)       19
	// Mn_Bus_Addr_Ln_2	varchar(38)
	// Mn_Bus_Sbrb	varchar(46)
	// Mn_Bus_Stt	varchar(3)
	// Mn_Bus_Pc	varchar(12)
	// Mn_Bus_Cntry_Cd	varchar(3)
	// Mn_Bus_DPID	integer               25

	// Ent_Eml	varchar(200)              26

	// Prty_Id_Blnk	varchar(1)
	// GST_Regn_Dt	varchar(8)
	// GST_Cancn_Dt	varchar(8)

	// Mn_Indy_Clsn	varchar(5)             30
	// Mn_Indy_Clsn_Descn	varchar(100)   31

	// ACN	varchar(20)
	// Sprsn_Ind	varchar(1)

	if len(r) < 3 {
		return Record{}
	}

	return Record{
		OrgNameHash:     hashSet(r[3:4]),
		OrgName:         r[3:4],
		NameHash:        hashSet(r[4:9]),
		Name:            r[4:9],
		TradingNameHash: hashSet(r[11:12]),
		TradingName:     r[11:12],
		SONAddressHash:  hashSet(r[12:19]),
		SONAddress:      r[12:19],
		BusAddressHash:  hashSet(r[19:26]),
		BusAddress:      r[19:26],
		EmailHash:       hashSet(r[26:27]),
		Email:           r[26:27],
		IndustryHash:    hashSet(r[30:32]),
		Industry:        r[30:32],
	}

}
func hashSet(i []string) []byte {
	hasher := md5.New()
	return hasher.Sum([]byte(strings.Join(i, "\t")))
}

func Diff(one *os.File, two *os.File, updateFile *os.File, date string) {

	hashMap := map[int]Record{}

	scanner := bufio.NewScanner(one)
	for scanner.Scan() {
		record := strings.SplitN(scanner.Text(), "\t", 34)
		if len(record) < 34 {
			log.Print("Skipping short row ", scanner.Text())
			os.Exit(1)
			continue
		}

		//if len(record) != 34 {
		//	log.Print("Wrong number of cells: "+string(line), len(record))
		//	continue
		//}
		i, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		hashMap[i] = hashRecord(record)
		//if len(hashMap[i].Name)> 0 && hashMap[i].Name[0] != "" {
		//	fmt.Print(len(hashMap[i].Name));
		//	os.Exit(1);
		//}
	}

	scanner = bufio.NewScanner(two)

	updated := 0
	newRecords := 0

	for scanner.Scan() {
		record := strings.SplitN(scanner.Text(), "\t", 34)
		if len(record) < 34 {
			log.Print("Skipping short row ", scanner.Text())
			os.Exit(1)
			continue
		}

		//if len(record) != 34 {
		//	log.Print("Wrong number of cells: "+string(line), len(record))
		//	continue
		//}
		i, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		if old, ok := hashMap[i]; ok {
			nr := hashRecord(record)

			combined := []string{
				record[0],
				record[1],
				record[2],
			}
			update := 0
			if !bytes.Equal(old.OrgNameHash, nr.OrgNameHash) {
				combined = append(combined, nr.OrgName...)
				combined = append(combined, old.OrgName...)
				combined = append(combined, date)
				update = 1
			} else {
				combined = append(combined, old.OrgName...)
				combined = append(combined, "")
				combined = append(combined, date)
			}
			if !bytes.Equal(old.NameHash, nr.NameHash) {
				combined = append(combined, nr.Name...)
				combined = append(combined, old.Name...)
				combined = append(combined, date)
				update = 1
			} else {
				combined = append(combined, old.Name...)
				combined = append(combined, "", "", "", "", "", "")
			}
			combined = append(combined,
				record[9],
				record[10])

			if !bytes.Equal(old.TradingNameHash, nr.TradingNameHash) {
				combined = append(combined, nr.TradingName...)
				combined = append(combined, old.TradingName...)
				combined = append(combined, date)
				update = 1
			} else {
				combined = append(combined, old.TradingName...)
				combined = append(combined, "")
				combined = append(combined, date)
			}
			if !bytes.Equal(old.SONAddressHash, nr.SONAddressHash) {
				combined = append(combined, nr.SONAddress...)
				combined = append(combined, old.SONAddress...)
				combined = append(combined, date)
				update = 1
			} else {
				combined = append(combined, old.SONAddress...)
				combined = append(combined, "", "", "", "", "", "", "", "")
			}
			if !bytes.Equal(old.BusAddressHash, nr.BusAddressHash) {
				combined = append(combined, nr.BusAddress...)
				combined = append(combined, old.BusAddress...)
				combined = append(combined, date)
				update = 1
			} else {
				combined = append(combined, old.BusAddress...)
				combined = append(combined, "", "", "", "", "", "", "", "")
			}
			if !bytes.Equal(old.EmailHash, nr.EmailHash) {
				combined = append(combined, nr.Email...)
				combined = append(combined, old.Email...)
				combined = append(combined, date)
				update = 1
			} else {
				combined = append(combined, old.Email...)
				combined = append(combined, "")
				combined = append(combined, date)
			}
			combined = append(combined,
				record[27],
				record[28],
				record[29],
			)
			if !bytes.Equal(old.IndustryHash, nr.IndustryHash) {
				combined = append(combined, nr.Industry...)
				combined = append(combined, old.Industry...)
				combined = append(combined, date)
				update = 1
			} else {
				combined = append(combined, old.Industry...)
				combined = append(combined, "")
				combined = append(combined, date)
			}
			combined = append(combined,
				record[32],
				record[33],
			)
			updated += update
			fmt.Fprintln(updateFile, strings.Join(combined, "\t"))

		} else {
			if len(record) < 34 {
				log.Print("Skipping short row ", scanner.Text())
				os.Exit(1)
				continue
			}

			combined := []string{
				record[0],
				record[1],
				record[2],
				record[3], // Org
				"",        // Prev_org
				"",        // org_dt
				record[4],
				record[5],
				record[6],
				record[7],
				record[8],
				"", "", "", "", "", "", // prev name and date
				record[9],
				record[10],
				record[11],
				"", "", // Prev trading and date
				record[12],
				record[13],
				record[14],
				record[15],
				record[16],
				record[17],
				record[18],
				"", "", "", "", "", "", "", "", // prev SON and date
				record[19],
				record[20],
				record[21],
				record[22],
				record[23],
				record[24],
				record[25],
				"", "", "", "", "", "", "", "", // prev Bus Addr and date
				record[26],
				"", "", // Prev email and date
				record[27],
				record[28],
				record[29],
				record[30],
				record[31],
				"", "", "", // Prev indy and date
				record[32],
				record[33],
			}
			fmt.Fprintln(updateFile, strings.Join(combined, "\t"))
			newRecords++
		}
	}

	fmt.Print("Updated Records: %d\nNew Records: %d", updated, newRecords)

}
