package main

import (
	"testing"

	"github.com/brpratt/advent-of-code/2019/intcode"
	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		program := intcode.FromText(lines[0])

		got := part01(program)
		expected := 4138658

		if got != expected {
			t.Fatalf("[expected %d] [actual %d]", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		program := intcode.FromText(lines[0])

		got := part02(program)
		expected := 7264

		if got != expected {
			t.Fatalf("[expected %d] [actual %d]", expected, got)
		}
	})
}
