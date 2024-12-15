package main

import (
	"bufio"
	"log"
	"os"
)

func Solve(path string) (total int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	pattern := []string{}
	for scanner.Scan() {
		row := scanner.Text()

		if row != "" {
			pattern = append(pattern, row)
			continue
		}

		total += solvePattern(pattern)

		pattern = []string{}
	}

	total += solvePattern(pattern)

	return total
}

func solvePattern(pattern []string) (count int) {
	for _, row := range pattern {
		println(row)
	}
	println()

	for idx := 0; idx < len(pattern[0])-1; idx++ {
		if checkVerticalMirror(pattern, idx) {
			return idx + 1
		}
	}

	for idx := 0; idx < len(pattern)-1; idx++ {
		if checkHorizontalMirror(pattern, idx) {
			return (idx + 1) * 100
		}
	}

	return count
}

func checkVerticalMirror(pattern []string, mirrorIdx int) bool {
	for _, row := range pattern {
		if !checkRow(row, mirrorIdx) {
			return false
		}
	}
	return true
}

func checkRow(row string, mirrorIdx int) bool {
	for leftIdx := 0; leftIdx <= mirrorIdx; leftIdx++ {
		rightIdx := leftIdx + (mirrorIdx-leftIdx)*2 + 1
		if rightIdx >= len(row) {
			continue
		}
		if row[leftIdx] != row[rightIdx] {
			return false
		}

	}

	return true
}

func checkHorizontalMirror(pattern []string, mirrorIdx int) bool {
	for topIdx := 0; topIdx <= mirrorIdx; topIdx++ {
		botIdx := topIdx + (mirrorIdx-topIdx)*2 + 1
		if botIdx >= len(pattern) {
			continue
		}
		if pattern[topIdx] != pattern[botIdx] {
			return false
		}

	}

	return true
}
