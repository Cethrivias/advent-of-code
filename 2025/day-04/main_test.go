package day04

import (
	"bufio"
	"os"
	"testing"
)

func TestPart1Example(t *testing.T) {
	lines, err := readFile("example")
	if err != nil {
		t.Error(err)
	}

	expected := 13
	actual, err := part1(lines)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Expected Part 1 Example to return %d, but got %d", expected, actual)
	}
}

func TestPart1Input(t *testing.T) {
	lines, err := readFile("input")
	if err != nil {
		t.Error(err)
	}

	expected := 1409
	actual, err := part1(lines)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Expected Part 1 Input to return %d, but got %d", expected, actual)
	}
}

func TestPart2Example(t *testing.T) {
	lines, err := readFile("example")
	if err != nil {
		t.Error(err)
	}

	expected := 43
	actual, err := part2(lines)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Expected Part 2 Example to return %d, but got %d", expected, actual)
	}
}

func TestPart2Input(t *testing.T) {
	lines, err := readFile("input")
	if err != nil {
		t.Error(err)
	}

	expected := 8366
	actual, err := part2(lines)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("Expected Part 2 Input to return %d, but got %d", expected, actual)
	}
}

func BenchmarkPart2Input(b *testing.B) {
	lines, err := readFile("input")
	if err != nil {
		b.Error(err)
	}

	for b.Loop() {
		part2(lines)
	}
}

func readFile(filename string) (lines [][]rune, err error) {
	lines = [][]rune{}

	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, []rune(scanner.Text()))
	}

	return lines, nil
}
