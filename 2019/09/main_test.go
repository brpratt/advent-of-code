package main

import (
	"testing"

	"github.com/brpratt/advent-of-code/2019/intcode"
	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	lines := file.Must(file.ReadLines("input.txt"))
	program := intcode.FromText(lines[0])

	expected := 4080871669
	got := part01(program)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestPart02(t *testing.T) {
	lines := file.Must(file.ReadLines("input.txt"))
	program := intcode.FromText(lines[0])

	expected := 75202
	got := part02(program)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
