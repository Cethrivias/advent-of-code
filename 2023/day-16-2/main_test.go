package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 51 {
		t.Errorf("Expected 51, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 7716 {
		t.Errorf("Expected 7716, but got %d", res)
	}
}
