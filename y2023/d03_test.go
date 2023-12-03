package y2023_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2023"
)

func TestSolveD03P01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

		expected := "4361"

		got, err := y2023.SolveD03P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}

	})
	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d03.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD03P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "533775"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD03P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

		expected := "467835"

		answer, err := y2023.SolveD03P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if answer != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, answer)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d03.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD03P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "78236071"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
