package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
		rounds := parseHandRounds(strings.NewReader(input))

		expected := 6440
		got := part01(rounds)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		rounds := parseHandRounds(input)

		expected := 253954294
		got := part01(rounds)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
		rounds := parseHandRounds(strings.NewReader(input))

		expected := 5905
		got := part02(rounds)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		rounds := parseHandRounds(input)

		expected := 254837398
		got := part02(rounds)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}
