package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(path string) int {
	total := 0

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		readingsStr := strings.Split(text, " ")
		readings := make([]int, len(readingsStr))

		for i, readingStr := range readingsStr {
			reading, err := strconv.Atoi(readingStr)
			if err != nil {
				log.Fatal(err)
			}
			readings[i] = reading
		}

		result := extrapolate(readings)
		total += result
	}

	return total
}

func extrapolate(readings []int) (result int) {
	done := true
	for _, reading := range readings {
		if reading != 0 {
			done = false
			break
		}
	}
	if done {
		return 0
	}

	diffs := []int{}

	for i := 1; i < len(readings); i++ {
		diffs = append(diffs, readings[i]-readings[i-1])
	}

	return readings[len(readings)-1] + extrapolate(diffs)
}
