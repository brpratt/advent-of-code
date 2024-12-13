package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		input := `AAAA
BBCD
BBCC
EEEC`

		expected := 140
		got := part01(parseInput(strings.NewReader(input)))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

		expected := 772
		got := part01(parseInput(strings.NewReader(input)))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("example 3", func(t *testing.T) {
		input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

		expected := 1930
		got := part01(parseInput(strings.NewReader(input)))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		expected := 1344578
		got := part01(parseInput(input))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}
