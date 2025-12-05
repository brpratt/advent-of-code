package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `3-5
10-14
16-20
12-18

1
5
8
11
17
32
`

		idranges, ids := parseInput(strings.NewReader(input))

		expected := 3
		got := part01(idranges, ids)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		idranges, ids := parseInput(input)

		expected := 598
		got := part01(idranges, ids)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}
