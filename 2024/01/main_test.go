package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `3   4
4   3
2   5
1   3
3   9
3   3`
		left, right := parseLists(strings.NewReader(input))

		expected := 11
		got := part01(left, right)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		left, right := parseLists(input)

		expected := 1651298
		got := part01(left, right)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `3   4
4   3
2   5
1   3
3   9
3   3`
		left, right := parseLists(strings.NewReader(input))

		expected := 31
		got := part02(left, right)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		left, right := parseLists(input)

		expected := 21306195
		got := part02(left, right)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}
