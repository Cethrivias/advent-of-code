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

	if res != 13663968099527 {
		t.Errorf("Expected to get 13663968099527, but got %d", res)
	}
}
