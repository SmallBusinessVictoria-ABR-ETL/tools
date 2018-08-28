package sbv_abr_etl

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"log"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	hash []byte
}

func Diff(one *os.File, two *os.File) {

	hasher := md5.New()
	hashMap := map[int]Record{}

	scanner := bufio.NewScanner(one)
	for scanner.Scan() {
		record := strings.SplitN(scanner.Text(), "\t", 2)

		//if len(record) != 34 {
		//	log.Print("Wrong number of cells: "+string(line), len(record))
		//	continue
		//}
		i, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		hashMap[i] = Record{hasher.Sum([]byte(scanner.Text()))}
	}

	scanner = bufio.NewScanner(two)
	for scanner.Scan() {
		record := strings.SplitN(scanner.Text(), "\t", 2)

		//if len(record) != 34 {
		//	log.Print("Wrong number of cells: "+string(line), len(record))
		//	continue
		//}
		i, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		if hash, ok := hashMap[i]; ok {
			if !bytes.Equal(hash.hash, hasher.Sum([]byte(scanner.Text()))) {

				log.Print("Updated Record")
				//fmt.Println(string(record[0]), string(line))
				//fmt.Println(string(hash.record[0]), string(hash.line))

			} else {
				//delete(hashMap, string(record[0]))
			}
		} else {
			log.Print("New Record: " + scanner.Text())
		}
	}

}
