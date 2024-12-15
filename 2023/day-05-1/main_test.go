package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 35 {
		t.Errorf("Result should be 35, but got %d", res)
	}
}

func TestSolve(t *testing.T) {
	res := Solve("input.txt")

	if res != 261668924 {
		t.Errorf("Result should be 261668924, but got %d", res)
	}
}
