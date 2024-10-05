package main

import (
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		masses := []int{12, 14, 1969, 100756}

		expected := 34241
		got := part01(masses)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		masses := file.Must(file.ReadNumbers("input.txt"))

		expected := 3317668
		got := part01(masses)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}

func TestSolveD01P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		masses := []int{14, 1969, 100756}

		expected := 51314
		got := part02(masses)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		masses := file.Must(file.ReadNumbers("input.txt"))

		expected := 4973628
		got := part02(masses)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}
