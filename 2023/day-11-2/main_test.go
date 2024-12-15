package main

import "testing"

func TestSolveExample2(t *testing.T) {
	res := Solve("input_test.txt", 2)

	if res != 374 {
		t.Errorf("Expected 374, but got %d", res)
	}
}

func TestSolveExample10(t *testing.T) {
	res := Solve("input_test.txt", 10)

	if res != 1030 {
		t.Errorf("Expected 1030, but got %d", res)
	}
}

func TestSolveExample100(t *testing.T) {
	res := Solve("input_test.txt", 100)

	if res != 8410 {
		t.Errorf("Expected 8410, but got %d", res)
	}
}

func TestSolveTask(t *testing.T) {
	res := Solve("input.txt", 1_000_000)

	if res != 622120986954 {
		t.Errorf("Expected 622120986954, but got %d", res)
	}
}
