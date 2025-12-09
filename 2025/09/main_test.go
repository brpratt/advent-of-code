package main

import (
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

		lines := strings.Split(input, "\n")
		points := parsePoints(lines)

		expected := 50
		got := part01(points)

		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		points := parsePoints(lines)

		expected := 4763040296
		got := part01(points)

		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}
