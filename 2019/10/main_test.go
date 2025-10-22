package main

import (
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/file"
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

func TestPart01(t *testing.T) {
	t.Run("examples", func(t *testing.T) {

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
				8,
			},
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
				33,
			},
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
				35,
			},
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
				41,
			},
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
				210,
			},
		}

		for _, test := range tests {
			lines := strings.Fields(test.input)
			result := part01(lines)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))

		expected := 326
		got := part01(lines)

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	t.Run("examples", func(t *testing.T) {
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
				802,
			},
		}

		for _, test := range tests {
			lines := strings.Fields(test.input)
			result := part02(lines, test.count)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))

		expected := 1623
		got := part02(lines, 200)

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}
