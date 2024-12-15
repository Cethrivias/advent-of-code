package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 46 {
		t.Errorf("Result should be 46, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 24261545 {
		t.Errorf("Result should be 24261545, but got %d", res)
	}
}
