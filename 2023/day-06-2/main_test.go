package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")
	if res != 71503 {
		t.Errorf("Result should be 71503, but found %d", res)
	}
}

func TestSolveInput(t *testing.T) {
	res := Solve("input.txt")
	if res != 40087680 {
		t.Errorf("Result should be 40087680, but found %d", res)
	}
}
