package day08

import (
	"bufio"
	"os"
	"testing"
)

func TestDayPart1(t *testing.T) {
	tests := []struct {
		file        string
		connections int
		expected    int
	}{
		{
			file:        "example",
			connections: 10,
			expected:    40,
		},
		{
			file:        "input",
			connections: 1000,
			expected:    54180,
		},
	}

	for _, test := range tests {
		lines, err := readFile(test.file)
		if err != nil {
			t.Error(err)
		}

		actual, err := part1(lines, test.connections)
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
			expected: 25272,
		},
		{
			file:     "input",
			expected: 25325968,
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

func readFile(filename string) (lines []string, err error) {
	lines = []string{}

	file, err := os.Open(filename)
	if err != nil {
		return lines, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
