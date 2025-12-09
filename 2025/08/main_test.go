package main

import (
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

		lines := strings.Split(input, "\n")
		boxes := parseBoxes(lines)

		expected := 40
		got := part01(boxes, 10)

		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		boxes := parseBoxes(lines)

		expected := 117000
		got := part01(boxes, 1000)

		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`

		lines := strings.Split(input, "\n")
		boxes := parseBoxes(lines)

		expected := 25272
		got := part02(boxes)

		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		boxes := parseBoxes(lines)

		expected := 8368033065
		got := part02(boxes)

		if got != expected {
			t.Fatalf("expected %d, got %d", expected, got)
		}
	})
}
