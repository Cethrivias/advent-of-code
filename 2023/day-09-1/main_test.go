package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 114 {
		t.Errorf("Expected 114, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 1731106378 {
		t.Errorf("Expected 1731106378, but got %d", res)
	}
}
