package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 21 {
		t.Errorf("Expected 21, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 7490 {
		t.Errorf("Expected 7490, but got %d", res)
	}
}

func TestValidateRecords(t *testing.T) {
	tests := []ValidateRecordsData{
		{
			records:  "####.#...#...",
			checksum: []int{4, 1, 1},
			result:   true,
		},
		{
			records:  "####.#......#",
			checksum: []int{4, 1, 1},
			result:   true,
		},
		{
			records:  ".###......##",
			checksum: []int{3, 2, 1},
			result:   false,
		},
		{
			records:  "############",
			checksum: []int{3, 2, 1},
			result:   false,
		},
	}

	for _, test := range tests {
		res := validateRecords(test.records, test.checksum)

		if res != test.result {
			t.Errorf("'%s' %+v should be %t, but got %t", test.records, test.checksum, test.result, res)
		}

	}
}

type ValidateRecordsData struct {
	records  string
	checksum []int
	result   bool
}

func TestCountCombinations(t *testing.T) {
	res, ok := countCombinations("?###????????", []int{3, 2, 1})

	if !ok || res != 10 {
		t.Errorf("Expected 10, but got %d", res)
	}
}
