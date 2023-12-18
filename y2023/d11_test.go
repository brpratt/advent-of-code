package y2023_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2023"
)

func TestSolveD11P01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

		expected := "374"

		got, err := y2023.SolveD11P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d11.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD11P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "9536038"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD11P02(t *testing.T) {
	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d11.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD11P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "447744640566"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
