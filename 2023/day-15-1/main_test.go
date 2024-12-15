package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 1320 {
		t.Errorf("Expected 1320, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 517315 {
		t.Errorf("Expected 517315, but got %d", res)
	}
}
