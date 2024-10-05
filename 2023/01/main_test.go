package main

import (
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestSolveD01P01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
		lines := strings.Split(input, "\n")

		expected := 142
		got := part01(lines)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))

		expected := 55123
		got := part01(lines)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}

func TestSolveD01P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
		lines := strings.Split(input, "\n")

		expected := 281
		got := part02(lines)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))

		expected := 55260
		got := part02(lines)

		if got != expected {
			t.Fatalf("[expected %d] [got %d]", expected, got)
		}
	})
}
