package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")
	if res != 136 {
		t.Errorf("Expected 136, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")
	if res != 108826 {
		t.Errorf("Expected 108826, but got %d", res)
	}
}
