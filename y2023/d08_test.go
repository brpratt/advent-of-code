package y2023_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2023"
)

func TestSolveD08P01(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		input := `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

		expected := "2"

		got, err := y2023.SolveD08P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		input := `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

		expected := "6"

		got, err := y2023.SolveD08P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d08.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD08P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "18113"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD08P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

		expected := "6"

		got, err := y2023.SolveD08P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d08.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD08P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "12315788159977"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
