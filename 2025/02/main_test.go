package main

import (
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
		idranges := parseRanges(input)

		expected := 1227775554
		got := part01(idranges)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(file.ReadLines("input.txt"))[0]
		idranges := parseRanges(input)

		expected := 24747430309
		got := part01(idranges)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
		idranges := parseRanges(input)

		expected := 4174379265
		got := part02(idranges)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(file.ReadLines("input.txt"))[0]
		idranges := parseRanges(input)

		expected := 30962646823
		got := part02(idranges)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}
