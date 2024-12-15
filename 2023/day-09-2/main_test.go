package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 2 {
		t.Errorf("Expected 2, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 1087 {
		t.Errorf("Expected 1087, but got %d", res)
	}
}
