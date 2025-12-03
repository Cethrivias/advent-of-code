package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	for _, line := range lines {
		fmt.Println(line)
	}
}

func part2(lines []string) {
	// noop
}
