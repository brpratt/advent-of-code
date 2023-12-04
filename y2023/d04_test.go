package y2023_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2023"
)

func TestSolveD04P01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

		expected := "13"

		got, err := y2023.SolveD04P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}

	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d04.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD04P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "18519"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD04P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

		expected := "30"

		answer, err := y2023.SolveD04P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if answer != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, answer)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d04.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD04P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "11787590"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
