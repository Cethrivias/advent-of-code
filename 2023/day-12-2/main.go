package main

import (
	"bufio"
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

		tmp := strings.Split(text, " ")

		records := tmp[0]
		checksumStr := tmp[1]
		for i := 0; i < 4; i++ {
			records += "?" + tmp[0]
		}
		for i := 0; i < 4; i++ {
			checksumStr += "," + tmp[1]
		}

		checksum := make([]int, 0)
		for _, val := range strings.Split(checksumStr, ",") {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}

			checksum = append(checksum, valInt)
		}

		total += countCombinations(records, checksum, 0)
	}

	return total
}

func countCombinations(records string, checksum []int, segmentLength int) (count int) {
	if len(records) < sum(checksum)-segmentLength+(len(checksum)-1) {
		return 0
	}

	if 0 == len(records) {
		if segmentLength == 0 && len(checksum) > 0 {
			return 0
		}
		if segmentLength > 0 && (len(checksum) == 0 || checksum[0] != segmentLength) {
			return 0
		}
		if len(checksum) > 1 {
			return 0
		}

		return 1
	}

	switch records[0] {
	case '#':
		segmentLength++

		if len(checksum) == 0 || checksum[0] < segmentLength {
			return 0
		}

		count += Memo(records[1:], checksum, segmentLength, countCombinations)
	case '.':
		if segmentLength != 0 {
			if checksum[0] != segmentLength {
				return 0
			}
			segmentLength = 0

			checksum = checksum[1:]
		}
		count += Memo(records[1:], checksum, segmentLength, countCombinations)
	case '?':
		for _, state := range []string{".", "#"} {
			if state == "." && segmentLength > 0 && checksum[0] > segmentLength {
				continue
			}

			records = state + records[1:]
			count += Memo(records, checksum, segmentLength, countCombinations)
		}
	}

	return count
}
