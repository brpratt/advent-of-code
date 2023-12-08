package y2023_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2023"
)

func TestSolveD07P01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

		expected := "6440"

		got, err := y2023.SolveD07P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d07.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD07P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "253954294"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD07P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

		expected := "5905"

		got, err := y2023.SolveD07P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d07.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD07P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "254837398"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
