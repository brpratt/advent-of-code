package day11

import (
	"os"
	"testing"
)

func TestSolvePart01(t *testing.T) {
	tests := []struct {
		program  []int
		expected int
	}{
		{
			[]int{3, 0, 104, 1, 99},
			1,
		},
	}

	for _, test := range tests {
		result := SolvePart01(test.program)
		if result != test.expected {
			t.Errorf("expected %d, got %d", test.expected, result)
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{1, 1967},
		{2, 0},
	}

	for _, test := range tests {
		file, err := os.Open("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		got := Solve(test.part, file)
		if got != test.expected {
			t.Errorf("failed part %d: expected %d, got %d", test.part, test.expected, got)
		}

		file.Close()
	}
}
