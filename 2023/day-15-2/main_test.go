package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 145 {
		t.Errorf("Expected 145, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 247763 {
		t.Errorf("Expected 247763, but got %d", res)
	}
}
