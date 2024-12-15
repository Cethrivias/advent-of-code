package main

import "testing"

func TestSolveExample(t *testing.T) {
    res := Solve("input_test.txt")

    if res != 102 {
        t.Errorf("Expected 102, but got %d", res)
    }
}
