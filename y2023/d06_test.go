package y2023_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2023"
)

func TestSolveD06P01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `Time:      7  15   30
Distance:  9  40  200`

		expected := "288"

		got, err := y2023.SolveD06P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d06.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD06P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "1108800"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD06P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `Time:      7  15   30
Distance:  9  40  200`

		expected := "288"

		got, err := y2023.SolveD06P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d06.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD06P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "36919753"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
