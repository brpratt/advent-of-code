package main

import (
	"slices"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		digits := []int{1, 1, 2, 2, 5, 6, 7, 8, 9, 0, 1, 2}
		got := part01(digits, 3, 2)
		expected := 4

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		digits := toDigits(lines[0])

		expected := 1452
		got := part01(digits, 25, 6)

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		digits := []int{0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0}
		got := part02(digits, 2, 2)
		expected := []int{0, 1, 1, 0}

		if !slices.Equal(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		digits := toDigits(lines[0])

		expected := []int{1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0}
		got := part02(digits, 25, 6)

		if !slices.Equal(expected, got) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}
