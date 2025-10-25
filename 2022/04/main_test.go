package main

import (
	"os"
	"testing"
)

func TestSolveExample(t *testing.T) {
	file, err := os.Open("test_input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	part1, part2 := solve(file)

	part1Expected := "2"
	part2Expected := "4"

	if part1 != part1Expected {
		t.Errorf("[Part 1]: expected %s, got %s", part1Expected, part1)
	}

	if part2 != part2Expected {
		t.Errorf("[Part 2]: expected %s, got %s", part2Expected, part2)
	}
}

func TestSolveActual(t *testing.T) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	part1, part2 := solve(file)

	part1Expected := "562"
	part2Expected := "924"

	if part1 != part1Expected {
		t.Errorf("[Part 1]: expected %s, got %s", part1Expected, part1)
	}

	if part2 != part2Expected {
		t.Errorf("[Part 2]: expected %s, got %s", part2Expected, part2)
	}
}
