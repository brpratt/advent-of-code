package main

import (
	"os"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

		rotations := parseRotations(strings.NewReader(input))

		expected := 3
		got := part01(rotations)

		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		rotations := parseRotations(input)

		expected := 1177
		got := part01(rotations)

		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

		rotations := parseRotations(strings.NewReader(input))

		expected := 6
		got := part02(rotations)

		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		input := file.Must(os.Open("input.txt"))
		defer input.Close()

		rotations := parseRotations(input)

		expected := 6768
		got := part02(rotations)

		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}
