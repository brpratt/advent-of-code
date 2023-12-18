package y2023_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2023"
)

func TestSolveD13P01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

		expected := "405"

		got, err := y2023.SolveD13P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d13.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD13P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "27505"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD13P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

		expected := "400"

		got, err := y2023.SolveD13P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d13.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD13P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := ""

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
