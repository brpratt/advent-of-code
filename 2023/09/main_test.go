package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
		report := parseReport(strings.NewReader(input))

		expected := 114
		got := part01(report)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		report := parseReport(input)

		expected := 1762065988
		got := part01(report)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
		report := parseReport(strings.NewReader(input))

		expected := 2
		got := part02(report)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		report := parseReport(input)

		expected := 1066
		got := part02(report)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}
