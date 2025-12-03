package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `987654321111111
811111111111119
234234234234278
818181911112111`

		banks := parseBanks(strings.NewReader(input))

		expected := 357
		got := part01(banks)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		banks := parseBanks(input)

		expected := 17085
		got := part01(banks)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `987654321111111
811111111111119
234234234234278
818181911112111`

		banks := parseBanks(strings.NewReader(input))

		expected := 3121910778619
		got := part02(banks)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		banks := parseBanks(input)

		expected := 169408143086082
		got := part02(banks)

		if expected != got {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}
