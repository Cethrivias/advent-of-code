package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 46 {
		t.Errorf("Expected 46, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 7472 {
		t.Errorf("Expected 7472, but got %d", res)
	}
}
