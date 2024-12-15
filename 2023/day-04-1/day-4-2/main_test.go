package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")

	if res != 30 {
		t.Errorf("Res should be 30, but got %d", res)

	}
}

func TestSolve(t *testing.T) {
	res := Solve("input.txt")

	if res != 5422730 {
		t.Errorf("Res should be 5422730, but got %d", res)
	}
}
