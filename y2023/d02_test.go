package y2023_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2023"
)

func TestSolveD02P01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

		expected := "8"

		got, err := y2023.SolveD02P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}

	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d02.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD02P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "3099"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD02P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

		expected := "2286"

		answer, err := y2023.SolveD02P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if answer != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, answer)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d02.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD02P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "72970"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
