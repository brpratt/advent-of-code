package main

import (
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := []string{
			"COM)B",
			"B)C",
			"C)D",
			"D)E",
			"E)F",
			"B)G",
			"G)H",
			"D)I",
			"E)J",
			"J)K",
			"K)L",
		}

		result := part01(input)
		if result != 42 {
			t.Errorf("expected 42, got %d", result)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(file.ReadLines("input.txt"))

		expected := part01(input)
		got := 308790

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := []string{
			"COM)B",
			"B)C",
			"C)D",
			"D)E",
			"E)F",
			"B)G",
			"G)H",
			"D)I",
			"E)J",
			"J)K",
			"K)L",
			"K)YOU",
			"I)SAN",
		}

		result := part02(input)
		if result != 4 {
			t.Errorf("expected 4, got %d", result)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(file.ReadLines("input.txt"))

		expected := part02(input)
		got := 472

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}
