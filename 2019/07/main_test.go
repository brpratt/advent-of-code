package main

import (
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("examples", func(t *testing.T) {
		tests := []struct {
			program []int
			signal  int
		}{
			{[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, 43210},
			{[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, 54321},
			{[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, 65210},
		}

		for _, test := range tests {
			result := part01(test.program)
			if result != test.signal {
				t.Errorf("expected signal %d, got %d", test.signal, result)
			}
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		program := toIntcode(lines[0])

		expected := 116680
		got := part01(program)

		if got != expected {
			t.Errorf("expected signal %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("examples", func(t *testing.T) {
		tests := []struct {
			program []int
			signal  int
		}{
			{[]int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}, 139629729},
			{[]int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}, 18216},
		}

		for _, test := range tests {
			result := part02(test.program)
			if result != test.signal {
				t.Errorf("expected signal %d, got %d", test.signal, result)
			}
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		program := toIntcode(lines[0])

		expected := 89603079
		got := part02(program)

		if got != expected {
			t.Errorf("expected signal %d, got %d", expected, got)
		}
	})
}
