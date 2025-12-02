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

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	position := 50
	password := 0

	for _, line := range lines {
		direction := 1
		if line[0] == 'L' {
			direction = -1
		}
		turns, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		turns %= 100
		turns *= direction
		position += turns

		if position > 99 {
			position -= 100
		} else if position < 0 {
			position += 100
		}
		if position == 0 {
			password++
		}
	}

	fmt.Printf("Part 1: %d\n", password)
}

func part2(lines []string) {
	position := 50
	password := 0

	for _, line := range lines {
		direction := 1
		if line[0] == 'L' {
			direction = -1
		}
		turns, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		circles := turns / 100
		password += circles

		turns %= 100

		for range turns {
			position += direction
			if position > 99 {
				position -= 100
			} else if position < 0 {
				position += 100
			}
			if position == 0 {
				password++
			}
		}
	}

	fmt.Printf("Part 2: %d\n", password)
}
