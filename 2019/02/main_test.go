package main

import (
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		values := strings.Split(lines[0], ",")
		program := toProgram(values)

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
		values := strings.Split(lines[0], ",")
		program := toProgram(values)

		got := part02(program)
		expected := 7264

		if got != expected {
			t.Fatalf("[expected %d] [actual %d]", expected, got)
		}
	})
}
