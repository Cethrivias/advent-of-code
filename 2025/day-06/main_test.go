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
			expected: 4277556,
		},
		{
			file:     "input",
			expected: 5524274308182,
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
			expected: 3263827,
		},
		{
			file:     "input",
			expected: 8843673199391,
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
