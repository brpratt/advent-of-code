package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

		g := parseGrid(strings.NewReader(input))

		expected := 13
		got := part01(g)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		g := parseGrid(input)

		expected := 1486
		got := part01(g)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

		g := parseGrid(strings.NewReader(input))

		expected := 43
		got := part02(g)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		g := parseGrid(input)

		expected := 9024
		got := part02(g)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}
