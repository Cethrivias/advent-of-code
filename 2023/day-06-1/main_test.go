package main

import "testing"

func TestSolveExample(t *testing.T) {
	res := Solve("input_test.txt")
	if res != 288 {
		t.Errorf("Result should be 288, but found %d", res)
	}
}

func TestSolveInput(t *testing.T) {
	res := Solve("input.txt")
	if res != 1731600 {
		t.Errorf("Result should be 1731600, but found %d", res)
	}
}
