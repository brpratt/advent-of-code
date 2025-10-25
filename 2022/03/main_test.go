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

	part1Expected := "157"
	part2Expected := "70"

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

	part1Expected := "7850"
	part2Expected := "2581"

	if part1 != part1Expected {
		t.Errorf("[Part 1]: expected %s, got %s", part1Expected, part1)
	}

	if part2 != part2Expected {
		t.Errorf("[Part 2]: expected %s, got %s", part2Expected, part2)
	}
}
