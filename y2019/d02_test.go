package y2019_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/brpratt/advent-of-code/y2019"
)

func TestSolveD02P01(t *testing.T) {
	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d02.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2019.SolveD02P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "4138658"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD02P02(t *testing.T) {
	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d02.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2019.SolveD02P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "7264"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
