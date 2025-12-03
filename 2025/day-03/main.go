package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}

func part1(lines []string) int {
	totalPower := 0

	for _, line := range lines {
		valOne := '0'
		valTwo := '0'
		idxOne := 0

		for i := 0; i < len(line)-1; i++ {
			val := rune(line[i])
			if val <= valOne {
				continue
			}
			valOne = val
			idxOne = i
			if val == '9' {
				break
			}
		}

		for i := idxOne + 1; i < len(line); i++ {
			val := rune(line[i])
			if val <= valTwo {
				continue
			}
			valTwo = val
			if val == '9' {
				break
			}
		}

		bankJoltage, err := strconv.Atoi(fmt.Sprintf("%c%c", valOne, valTwo))
		if err != nil {
			log.Fatal(err)
		}

		totalPower += bankJoltage
	}

	return totalPower
}

func part2(lines []string) int {
	totalPower := 0

	for _, line := range lines {
		values := []rune{}
		remainingLength := 12

		curIdx := -1

		for remainingLength > 0 {
			curVal := '0'
			for i := curIdx + 1; i < len(line)-remainingLength+1; i++ {
				val := rune(line[i])
				if val <= curVal {
					continue
				}
				curVal = val
				curIdx = i
				if val == '9' {
					break
				}
			}
			values = append(values, curVal)
			remainingLength--
		}

		bankJoltage, err := strconv.Atoi(string(values))
		if err != nil {
			log.Fatal(err)
		}

		totalPower += bankJoltage
	}

	return totalPower
}
