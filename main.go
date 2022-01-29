package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func csvReaderAndWriter() {
	start := time.Now()

	// 1. Open the file
	csvFile, err1 := os.Open("./test.csv")
	if err1 != nil {
		fmt.Println("An error encountered ::", err1)
		return
	}

	// 2. Initialize the reader
	reader := csv.NewReader(csvFile)

	// 3. Read all the records
	records, err2 := reader.ReadAll()
	if err2 != nil {
		fmt.Println("An error encountered ::", err2)
		return
	}

	// 4. Processing
	var newRecords [][]string
	for _, record := range records {
		duplicate := false
		indexToBeUpdated := -1
		toBeAddedLoan := 0

		for i, newRecord := range newRecords {
			if newRecord[0] == record[0] && newRecord[1] == record[1] {
				duplicate = true
				indexToBeUpdated = i
				toBeAddedLoan, _ = strconv.Atoi(newRecord[2])
			}
		}

		if !duplicate {
			newRecords = append(newRecords, record)
		} else {
			prevLoan, _ := strconv.Atoi(record[2])
			newRecords[indexToBeUpdated] = []string{record[0], record[1], strconv.Itoa(prevLoan + toBeAddedLoan)}
		}
	}

	// 5. Create a new file
	resultFile, err3 := os.Create("./result.csv")
	if err3 != nil {
		fmt.Println("An error encountered ::", err3)
		return
	}

	// 6. Initialize the writer
	writer := csv.NewWriter(resultFile)

	// 7. Write all the records
	err4 := writer.WriteAll(newRecords)
	if err4 != nil {
		fmt.Println("An error encountered ::", err4)
	}

	fmt.Println(time.Since(start))
}

func main() {
	csvReaderAndWriter()
}
