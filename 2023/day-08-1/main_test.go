package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 6 {
		t.Errorf("Expected to get 6, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 19199 {
		t.Errorf("Expected to get 19199, but got %d", res)
	}
}
