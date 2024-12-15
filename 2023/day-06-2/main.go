package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solve(path string) int {
	total := 0

	input := getInput(path)

	for x := 1; x < input.time; x++ {
		if (input.time-x)*x > input.distance {
			total++
		}
	}

	return total
}

func getInput(path string) Input {
	input := Input{}

	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	timeStr := ""
	for _, val := range strings.Split(strings.Split(scanner.Text(), ":")[1], " ") {
		if val == "" {
			continue
		}
		timeStr += val
	}
	time, _ := strconv.Atoi(timeStr)
	input.time = time

	scanner.Scan()
	distanceStr := ""
	for _, val := range strings.Split(strings.Split(scanner.Text(), ":")[1], " ") {
		if val == "" {
			continue
		}
		distanceStr += val
	}
	distance, _ := strconv.Atoi(distanceStr)
	input.distance = distance

	return input
}

type Input struct {
	time     int
	distance int
}
