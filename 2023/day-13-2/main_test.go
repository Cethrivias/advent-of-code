package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 400 {
		t.Errorf("Expected 400, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 33991 {
		t.Errorf("Expected 33991, but got %d", res)
	}
}
