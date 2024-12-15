package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solve(path string) int {
	total := 1

	input := getInput(path)

	for i, time := range input.times {
		distance := input.distances[i]
		wins := 0
		for x := 1; x < time; x++ {
			if (time-x)*x > distance {
				wins++
			}
		}
		if wins > 0 {
			total *= wins
		}
	}

	return total
}

func getInput(path string) Input {
	input := Input{}

	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	for _, val := range strings.Split(strings.Split(scanner.Text(), ":")[1], " ") {
		if val == "" {
			continue
		}
		time, _ := strconv.Atoi(val)
		input.times = append(input.times, time)
	}

	scanner.Scan()
	for _, val := range strings.Split(strings.Split(scanner.Text(), ":")[1], " ") {
		if val == "" {
			continue
		}
		distance, _ := strconv.Atoi(val)
		input.distances = append(input.distances, distance)
	}

	return input
}

type Input struct {
	times     []int
	distances []int
}
