package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 13 {
		t.Errorf("Res should be 13, but got %d", res)

	}
}

func TestSolve(t *testing.T) {
	res := Solve("input.txt")

	if res != 24733 {
		t.Errorf("Res should be 24733, but got %d", res)
	}
}
