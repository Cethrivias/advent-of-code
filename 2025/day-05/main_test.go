package day05

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
			expected: 3,
		},
		{
			file:     "input",
			expected: 868,
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
			expected: 14,
		},
		{
			file:     "input",
			expected: 354143734113772,
			// 355083055110630 -- too high
			// 261012033400408 -- too low
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
