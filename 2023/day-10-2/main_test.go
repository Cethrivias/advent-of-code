package main

import (
	"testing"
)

func TestSolveExample1(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 4 {
		t.Errorf("Expected 4, but got %d", res)
	}
}

func TestSolveExample2(t *testing.T) {
	res := Solve("input_test_2.txt")

	if res != 4 {
		t.Errorf("Expected 4, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt")

	if res != 287 {
		t.Errorf("Expected 287, but got %d", res)
	}
}
