package main

import (
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

		lines := strings.Split(input, "\n")
		blocks, operators := parseInput(lines)

		expected := 4277556
		got := part01(blocks, operators)

		if got != expected {
			t.Fatalf("expected %d, got %d", got, expected)
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		blocks, operators := parseInput(lines)

		expected := 4583860641327
		got := part01(blocks, operators)

		if got != expected {
			t.Fatalf("expected %d, got %d", got, expected)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

		lines := strings.Split(input, "\n")
		blocks, operators := parseInput(lines)

		expected := 3263827
		got := part02(blocks, operators)

		if got != expected {
			t.Fatalf("expected %d, got %d", got, expected)
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		blocks, operators := parseInput(lines)

		expected := 11602774058280
		got := part02(blocks, operators)

		if got != expected {
			t.Fatalf("expected %d, got %d", got, expected)
		}
	})
}
