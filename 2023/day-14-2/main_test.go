package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")
	if res != 64 {
		t.Errorf("Expected 64, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")
	if res != 99570 {
		t.Errorf("Expected 99570, but got %d", res)
	}
}
