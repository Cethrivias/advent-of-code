package day06

import (
	"bufio"
	"os"
	"testing"
)

func TestDayPart1(t *testing.T) {
	tests := []struct {
		file     string
		expected int
	}{
		{
			file:     "example",
			expected: 21,
		},
		{
			file:     "input",
			expected: 1646,
		},
	}

	for _, test := range tests {
		lines, err := readFile(test.file)
		if err != nil {
			t.Error(err)
		}

		actual, err := part1(lines)
		if err != nil {
			t.Error(err)
		}
		if actual != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, actual)
		}
	}
}

func TestDayPart2(t *testing.T) {
	tests := []struct {
		file     string
		expected int
	}{
		{
			file:     "example",
			expected: 40,
		},
		{
			file:     "input",
			expected: 32451134474991,
		},
	}

	for _, test := range tests {
		lines, err := readFile(test.file)
		if err != nil {
			t.Error(err)
		}

		actual, err := part2(lines)
		if err != nil {
			t.Error(err)
		}
		if actual != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, actual)
		}
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
