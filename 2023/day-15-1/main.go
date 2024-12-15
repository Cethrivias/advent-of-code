package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func Solve(path string) (total int) {
	input := getInput(path)

	for i := range input {
		total += hash(input[i])
	}

	return total
}

func hash(key string) (hash int) {
	for _, rune := range key {
		hash = ((hash + int(rune)) * 17) % 256
	}

	return hash
}

func getInput(path string) []string {
	input := []string{}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		input = strings.Split(text, ",")
	}

	return input
}
