package main

import (
	"testing"

	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	lines := file.Must(file.ReadLines("input.txt"))
	program := toIntcode(lines[0])

	got := part01(program)
	expected := 12428642

	if got != expected {
		t.Errorf("[expected %d] [got %d]", expected, got)
	}
}

func TestPart02(t *testing.T) {
	lines := file.Must(file.ReadLines("input.txt"))
	program := toIntcode(lines[0])

	got := part02(program)
	expected := 918655

	if got != expected {
		t.Errorf("[expected %d] [got %d]", expected, got)
	}
}
