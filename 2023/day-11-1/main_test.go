package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 374 {
		t.Errorf("Expected 374, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 10231178 {
		t.Errorf("Expected 10231178, but got %d", res)
	}
}
