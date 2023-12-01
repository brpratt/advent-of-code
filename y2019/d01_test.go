package y2019_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2019"
)

func TestSolveD01P01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `12
14
1969
100756`

		expected := "34241"

		got, err := y2019.SolveD01P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}

	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d01.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2019.SolveD01P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "3317668"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD01P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `14
1969
100756`

		expected := "51314"

		answer, err := y2019.SolveD01P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if answer != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, answer)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d01.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2019.SolveD01P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "4973628"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
