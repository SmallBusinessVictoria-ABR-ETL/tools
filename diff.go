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
	OrgNameHash         []byte
	OrgName             []string
	PrevOrgNameDate     []string
	NameHash            []byte
	Name                []string
	PrevNameDate        []string
	TradingNameHash     []byte
	TradingName         []string
	PrevTradingNameDate []string
	SONAddressHash      []byte
	SONAddress          []string
	PrevSONAddressDate  []string
	BusAddressHash      []byte
	BusAddress          []string
	PrevBusAddressDate  []string
	EmailHash           []byte
	Email               []string
	PrevEmailDate       []string
	IndustryHash        []byte
	Industry            []string
	PrevIndustryDate    []string
}

func hashRecord(r []string) Record {

	if len(r) < 3 {
		return Record{}
	}

	return Record{
		OrgNameHash:         hashSet(r[3:4]),
		OrgName:             r[3:4],
		PrevOrgNameDate:     []string{"", ""},
		NameHash:            hashSet(r[4:9]),
		Name:                r[4:9],
		PrevNameDate:        []string{"", "", "", "", "", ""},
		TradingNameHash:     hashSet(r[11:12]),
		TradingName:         r[11:12],
		PrevTradingNameDate: []string{"", ""},
		SONAddressHash:      hashSet(r[12:19]),
		SONAddress:          r[12:19],
		PrevSONAddressDate:  []string{"", "", "", "", "", "", "", ""},
		BusAddressHash:      hashSet(r[19:26]),
		BusAddress:          r[19:26],
		PrevBusAddressDate:  []string{"", "", "", "", "", "", "", ""},
		EmailHash:           hashSet(r[26:27]),
		Email:               r[26:27],
		PrevEmailDate:       []string{"", ""},
		IndustryHash:        hashSet(r[30:32]),
		Industry:            r[30:32],
		PrevIndustryDate:    []string{"", "", ""},
	}

}

/**
There are ones created by this tool
*/
func hashCombinedRecord(r []string) Record {
	if len(r) < 3 {
		return Record{}
	}

	return Record{
		OrgNameHash:     hashSet(r[3:4]),
		OrgName:         r[3:4],
		NameHash:        hashSet(r[6:11]),
		Name:            r[6:11],
		TradingNameHash: hashSet(r[19:20]),
		TradingName:     r[19:20],
		SONAddressHash:  hashSet(r[22:29]),
		SONAddress:      r[22:29],
		BusAddressHash:  hashSet(r[37:44]),
		BusAddress:      r[37:44],
		EmailHash:       hashSet(r[52:53]),
		Email:           r[52:53],
		IndustryHash:    hashSet(r[58:60]),
		Industry:        r[58:60],
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
		record := strings.SplitN(scanner.Text(), "\t", 64)
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

		if len(record) > 34 {
			hashMap[i] = hashCombinedRecord(record)
		} else {
			hashMap[i] = hashRecord(record)
		}
	}

	scanner = bufio.NewScanner(two)

	OrgNameChangeFile, _ := os.Create("OrgNameChange.txt")
	NameChangeFile, _ := os.Create("NameChange.txt")
	TradingNameChangeFile, _ := os.Create("TradingNameChange.txt")
	SONAddressChangeFile, _ := os.Create("SONAddressChange.txt")
	BusAddressChangeFile, _ := os.Create("BusAddressChange.txt")
	EmailChangeFile, _ := os.Create("EmailChange.txt")
	IndustryChangeFile, _ := os.Create("IndustryChange.txt")

	updated := 0
	newRecords := 0

	for scanner.Scan() {
		record := strings.SplitN(scanner.Text(), "\t", 34)
		if len(record) < 34 {
			log.Print("Skipping short row ", scanner.Text())
			os.Exit(1)
			continue
		}
		i, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		nr := hashRecord(record)

		if old, ok := hashMap[i]; ok {
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
				fmt.Fprintln(OrgNameChangeFile, record[0]+"\t"+strings.Join(nr.OrgName, "\t")+"\t"+date)
				update = 1
			} else {
				combined = append(combined, old.OrgName...)
				combined = append(combined, old.PrevOrgNameDate...)
			}
			if !bytes.Equal(old.NameHash, nr.NameHash) {
				combined = append(combined, nr.Name...)
				combined = append(combined, old.Name...)
				combined = append(combined, date)
				fmt.Fprintln(NameChangeFile, record[0]+"\t"+strings.Join(nr.Name, "\t")+"\t"+date)
				update = 1
			} else {
				combined = append(combined, old.Name...)
				combined = append(combined, old.PrevNameDate...)
			}
			combined = append(combined,
				record[9],
				record[10])

			if !bytes.Equal(old.TradingNameHash, nr.TradingNameHash) {
				combined = append(combined, nr.TradingName...)
				combined = append(combined, old.TradingName...)
				combined = append(combined, date)
				fmt.Fprintln(TradingNameChangeFile, record[0]+"\t"+strings.Join(nr.TradingName, "\t")+"\t"+date)
				update = 1
			} else {
				combined = append(combined, old.TradingName...)
				combined = append(combined, old.PrevTradingNameDate...)

			}
			if !bytes.Equal(old.SONAddressHash, nr.SONAddressHash) {
				combined = append(combined, nr.SONAddress...)
				combined = append(combined, old.SONAddress...)
				combined = append(combined, date)
				fmt.Fprintln(SONAddressChangeFile, record[0]+"\t"+strings.Join(nr.SONAddress, "\t")+"\t"+date)
				update = 1
			} else {
				combined = append(combined, old.SONAddress...)
				combined = append(combined, old.PrevSONAddressDate...)

			}
			if !bytes.Equal(old.BusAddressHash, nr.BusAddressHash) {
				combined = append(combined, nr.BusAddress...)
				combined = append(combined, old.BusAddress...)
				combined = append(combined, date)
				fmt.Fprintln(BusAddressChangeFile, record[0]+"\t"+strings.Join(nr.BusAddress, "\t")+"\t"+date)
				update = 1
			} else {
				combined = append(combined, old.BusAddress...)
				combined = append(combined, old.PrevBusAddressDate...)

			}
			if !bytes.Equal(old.EmailHash, nr.EmailHash) {
				combined = append(combined, nr.Email...)
				combined = append(combined, old.Email...)
				combined = append(combined, date)
				fmt.Fprintln(EmailChangeFile, record[0]+"\t"+strings.Join(nr.Email, "\t")+"\t"+date)
				update = 1
			} else {
				combined = append(combined, old.Email...)
				combined = append(combined, old.PrevEmailDate...)
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
				fmt.Fprintln(IndustryChangeFile, record[0]+"\t"+strings.Join(nr.Industry, "\t")+"\t"+date)
				update = 1
			} else {
				combined = append(combined, old.Industry...)
				combined = append(combined, old.PrevIndustryDate...)
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

			fmt.Fprintln(OrgNameChangeFile, record[0]+"\t"+strings.Join(nr.OrgName, "\t")+"\t"+date)
			fmt.Fprintln(NameChangeFile, record[0]+"\t"+strings.Join(nr.Name, "\t")+"\t"+date)
			fmt.Fprintln(TradingNameChangeFile, record[0]+"\t"+strings.Join(nr.TradingName, "\t")+"\t"+date)
			fmt.Fprintln(SONAddressChangeFile, record[0]+"\t"+strings.Join(nr.SONAddress, "\t")+"\t"+date)
			fmt.Fprintln(BusAddressChangeFile, record[0]+"\t"+strings.Join(nr.BusAddress, "\t")+"\t"+date)
			fmt.Fprintln(EmailChangeFile, record[0]+"\t"+strings.Join(nr.Email, "\t")+"\t"+date)
			fmt.Fprintln(IndustryChangeFile, record[0]+"\t"+strings.Join(nr.Industry, "\t")+"\t"+date)
		}
	}

	OrgNameChangeFile.Close()
	NameChangeFile.Close()
	TradingNameChangeFile.Close()
	SONAddressChangeFile.Close()
	BusAddressChangeFile.Close()
	EmailChangeFile.Close()
	IndustryChangeFile.Close()

	fmt.Printf("Updated Records: %d\nNew Records: %d", updated, newRecords)

}
