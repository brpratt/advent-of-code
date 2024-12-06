package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

		expected := 41
		got := part01(parseInput(strings.NewReader(input)))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		expected := 4789
		got := part01(parseInput(input))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

		expected := 6
		got := part02(parseInput(strings.NewReader(input)))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		expected := 1304
		got := part02(parseInput(input))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}
