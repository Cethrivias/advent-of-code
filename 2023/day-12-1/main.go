package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(path string) (total int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		print(text)

		tmp := strings.Split(text, " ")

		records := tmp[0]
		checksum := make([]int, 0)

		for _, val := range strings.Split(tmp[1], ",") {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}

			checksum = append(checksum, valInt)
		}

		count, _ := countCombinations(records, checksum)
		println(" => ", count)

		total += count
	}

	return total
}

func countCombinations(records string, checksum []int) (count int, ok bool) {
	idx := strings.Index(records, "?")
	if idx == -1 {
		if validateRecords(records, checksum) {
			fmt.Printf("records: %+v => %t\n", records, validateRecords(records, checksum))
			return 1, true
		}
		return count, false
	}

	for _, val := range []string{".", "#"} {
		recordsCurr := records[:idx] + val + records[idx+1:]

		if countCurr, ok := countCombinations(recordsCurr, checksum); ok {
			count += countCurr
		}
	}

	return count, true
}

func validateRecords(records string, checksum []int) bool {
	// fmt.Printf("Records: %+v\n", records)
	i := -1
	brokenCount := 0

	for _, val := range records {
		switch val {
		case '.':
			if brokenCount == 0 {
				continue
			}
			i++
			// fmt.Printf("validateRecords => checksum: %+v, idx: %d, brokenCount: %d, valid: %t\n", checksum, i, brokenCount, validateSegment(checksum, i, brokenCount))
			if !validateSegment(checksum, i, brokenCount) {
				return false
			}
			brokenCount = 0
		case '#':
			brokenCount++
		}
	}

	if brokenCount == 0 {
		return i == len(checksum)-1
	}

	i++
	return i == len(checksum)-1 && validateSegment(checksum, i, brokenCount)
}

func validateSegment(checksum []int, idx int, brokenCount int) bool {
	// fmt.Printf("validateSegment => checksum: %+v, idx: %d, brokenCount: %d\n", checksum, idx, brokenCount)
	if idx >= len(checksum) {
		return false
	}

	return checksum[idx] == brokenCount
}

type SpringState string

var (
	Operational SpringState = "."
	Damaged     SpringState = "#"
	Unknown     SpringState = "?"
)
