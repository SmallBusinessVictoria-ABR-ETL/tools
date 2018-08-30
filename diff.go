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
	hash []byte
}

func Diff(one *os.File, two *os.File, updates *os.File, newRecords *os.File) {

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

				fmt.Fprintln(updates, scanner.Text())

			} else {
				//delete(hashMap, string(record[0]))
			}
		} else {
			fmt.Fprintln(newRecords, scanner.Text())
		}
	}

}
