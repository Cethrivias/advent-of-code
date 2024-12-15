package main

import (
	"testing"
)

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 525152 {
		t.Errorf("Expected 525152, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 65607131946466 {
		t.Errorf("Expected 65607131946466, but got %d", res)
	}
}

func TestCountCombinations(t *testing.T) {
	// ???.### 1,1,3
	// .??..??...?##. 1,1,3
	// ?#?#?#?#?#?#?#? 1,3,1,6
	// ????.#...#... 4,1,1
	// ????.######..#####. 1,6,5
	// ?###???????? 3,2,1
	tests := []CountCombinationsData{
		{Records: "???.###", Checksum: []int{1, 1, 3}, Result: 1},
		{Records: ".??..??...?##.", Checksum: []int{1, 1, 3}, Result: 4},
		{Records: "?#?#?#?#?#?#?#?", Checksum: []int{1, 3, 1, 6}, Result: 1},
		{Records: "????.#...#...", Checksum: []int{4, 1, 1}, Result: 1},
		{Records: "????.######..#####.", Checksum: []int{1, 6, 5}, Result: 4},
		{Records: "?###????????", Checksum: []int{3, 2, 1}, Result: 10},
		{Records: "???.###????.###????.###????.###????.###", Checksum: []int{1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3}, Result: 1},
		{Records: "????.######..#####.?????.######..#####.?????.######..#####.?????.######..#####.?????.######..#####.", Checksum: []int{1, 6, 5, 1, 6, 5, 1, 6, 5, 1, 6, 5, 1, 6, 5}, Result: 2500},
		{
			Records:  "#?????????#?????#?????????#?????#?????????#?????#?????????#?????#?????????#????",
			Checksum: []int{1, 3, 2, 2, 1, 3, 2, 2, 1, 3, 2, 2, 1, 3, 2, 2, 1, 3, 2, 2},
			Result:   79434432,
		},
	}

	for _, test := range tests {
		res := countCombinations(test.Records, test.Checksum, 0)
		if res != test.Result {
			t.Errorf("%s %+v should have %d combinations, but got %d", test.Records, test.Checksum, test.Result, res)
		}
	}
}

type CountCombinationsData struct {
	Records  string
	Checksum []int
	Result   int
}

func TestCache(t *testing.T) {
	cache := NewCache()

	cache.Set("asdf", []int{1, 2, 3}, 4, 999)

	res, ok := cache.Get("asdf", []int{1, 2, 3}, 4)

	if !ok || res != 999 {
		t.Errorf("Expected true and 999, but got %t and %d", ok, res)
	}

	res, ok = cache.Get("qqqq", []int{4}, 0)

	if ok || res != 0 {
		t.Errorf("Expected false and 0, but got %t and %d", ok, res)
	}
}
