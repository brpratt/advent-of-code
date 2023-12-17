package y2023_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2023"
)

func TestSolveD10P01(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		input := `.....
.S-7.
.|.|.
.L-J.
.....`

		expected := "4"

		got, err := y2023.SolveD10P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		input := `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

		expected := "8"

		got, err := y2023.SolveD10P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d10.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD10P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "6842"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD10P02(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		input := `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

		expected := "4"

		got, err := y2023.SolveD10P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		input := `..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........`

		expected := "4"

		got, err := y2023.SolveD10P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
	t.Run("example 3", func(t *testing.T) {
		input := `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

		expected := "8"

		got, err := y2023.SolveD10P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("example 4", func(t *testing.T) {
		input := `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

		expected := "10"

		got, err := y2023.SolveD10P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
	/*
		t.Run("actual", func(t *testing.T) {
			file, err := os.Open(filepath.Join("input", "d10.txt"))
			if err != nil {
				t.Fatal(err)
			}

			defer file.Close()

			got, err := y2023.SolveD10P02(file)
			if err != nil {
				t.Fatal(err)
			}

			expected := "393"

			if got != expected {
				t.Fatalf("[expected %s] [actual %s]", expected, got)
			}
		})
	*/
}
