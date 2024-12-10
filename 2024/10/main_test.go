package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

		expected := 36
		got := part01(parseInput(strings.NewReader(input)))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		expected := 776
		got := part01(parseInput(input))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

		expected := 81
		got := part02(parseInput(strings.NewReader(input)))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		expected := 1657
		got := part02(parseInput(input))

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}
