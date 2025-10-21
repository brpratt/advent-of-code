package day10

import (
	"os"
	"strings"
	"testing"
)

func TestSlopeBetween(t *testing.T) {
	tests := []struct {
		a     point
		b     point
		slope slope
	}{
		{point{0, 0}, point{2, 7}, slope{7, 2}},
		{point{0, 0}, point{2, 4}, slope{2, 1}},
		{point{0, 0}, point{4, 6}, slope{3, 2}},
		{point{1, 1}, point{0, 0}, slope{-1, -1}},
	}

	for _, test := range tests {
		slope := slopeBetween(test.a, test.b)
		if slope != test.slope {
			t.Errorf("expected %v, got %v", test.slope, slope)
		}
	}
}

func TestSolvePart01(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`
.#..#
.....
#####
....#
...##`,
			8},
		{`
......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`,
			33},
		{`
#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`,
			35},
		{`
.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`,
			41},
		{`
.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
			210},
	}

	for _, test := range tests {
		lines := strings.Fields(test.input)
		result := SolvePart01(lines)
		if result != test.expected {
			t.Errorf("expected %d, got %d", test.expected, result)
		}
	}
}

func TestSolvePart02(t *testing.T) {
	tests := []struct {
		input    string
		count    int
		expected int
	}{
		{`
.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
			200,
			802},
	}

	for _, test := range tests {
		lines := strings.Fields(test.input)
		result := SolvePart02(lines, test.count)
		if result != test.expected {
			t.Errorf("expected %d, got %d", test.expected, result)
		}
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		part     int
		expected int
	}{
		{1, 326},
		{2, 1623},
	}

	for _, test := range tests {
		file, err := os.Open("input.txt")
		if err != nil {
			t.Fatal(err)
		}

		got := Solve(test.part, file)
		if got != test.expected {
			t.Errorf("failed part %d: expected %d, got %d", test.part, test.expected, got)
		}

		file.Close()
	}
}
