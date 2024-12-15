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

	pattern := [][]rune{}
	for scanner.Scan() {
		row := scanner.Text()

		if row != "" {
			pattern = append(pattern, []rune(row))
			continue
		}

		total += solvePattern(pattern)

		pattern = [][]rune{}
	}

	total += solvePattern(pattern)

	return total
}

func solvePattern(pattern [][]rune) (count int) {
	for _, row := range pattern {
		println(string(row))
	}
	println()

	mirrorIdx := -1
	isVertical := true
	for idx := 0; idx < len(pattern[0])-1; idx++ {
		if checkVerticalMirror(pattern, idx) {
			mirrorIdx = idx
			isVertical = true
			break
		}
	}

	if mirrorIdx == -1 {
		for idx := 0; idx < len(pattern)-1; idx++ {
			if checkHorizontalMirror(pattern, idx) {
				mirrorIdx = idx
				isVertical = false
				break
			}
		}
	}

	for i := range pattern {
		for j := range pattern[i] {
			pattern[i][j] = fixSmudge(pattern[i][j])

			for idx := 0; idx < len(pattern[0])-1; idx++ {
				if isVertical && idx == mirrorIdx {
					continue
				}
				if checkVerticalMirror(pattern, idx) {
					return idx + 1
				}
			}

			for idx := 0; idx < len(pattern)-1; idx++ {
				if !isVertical && idx == mirrorIdx {
					continue
				}
				if checkHorizontalMirror(pattern, idx) {
					return (idx + 1) * 100
				}
			}

			pattern[i][j] = fixSmudge(pattern[i][j])
		}
	}

	return count
}

func checkVerticalMirror(pattern [][]rune, mirrorIdx int) bool {
	for _, row := range pattern {
		if !checkRow(row, mirrorIdx) {
			return false
		}
	}
	return true
}

func checkRow(row []rune, mirrorIdx int) bool {
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

func checkHorizontalMirror(pattern [][]rune, mirrorIdx int) bool {
	for topIdx := 0; topIdx <= mirrorIdx; topIdx++ {
		botIdx := topIdx + (mirrorIdx-topIdx)*2 + 1
		if botIdx >= len(pattern) {
			continue
		}
		if !equalRows(pattern[topIdx], pattern[botIdx]) {
			return false
		}

	}

	return true
}

func equalRows(a []rune, b []rune) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func fixSmudge(r rune) rune {
	if r == '.' {
		return '#'
	} else {
		return '.'
	}
}
