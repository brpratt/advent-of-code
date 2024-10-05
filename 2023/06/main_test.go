package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `Time:      7  15   30
Distance:  9  40  200`
		times, distances := parseRaceDoc(strings.NewReader(input))

		expected := 288
		got := part01(times, distances)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		times, distances := parseRaceDoc(input)

		expected := 1108800
		got := part01(times, distances)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `Time:      7  15   30
Distance:  9  40  200`
		times, distances := parseRaceDoc(strings.NewReader(input))

		expected := 71503
		got := part02(times, distances)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		times, distances := parseRaceDoc(input)

		expected := 36919753
		got := part02(times, distances)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}
