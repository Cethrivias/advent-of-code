package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 405 {
		t.Errorf("Expected 405, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 27664 {
		t.Errorf("Expected 27664, but got %d", res)
	}
}
