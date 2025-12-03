package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func BenchmarkPart2(b *testing.B) {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for b.Loop() {
		part2(lines)
	}
}
