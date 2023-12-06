package y2023_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/brpratt/advent-of-code/y2023"
)

func TestSolveD05P01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

		expected := "35"

		got, err := y2023.SolveD05P01(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}

	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d05.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD05P01(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "31599214"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}

func TestSolveD05P02(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

		expected := "46"

		got, err := y2023.SolveD05P02(strings.NewReader(input))
		if err != nil {
			t.Fatal(err)
		}
		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}

	})

	t.Run("actual", func(t *testing.T) {
		file, err := os.Open(filepath.Join("input", "d05.txt"))
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		got, err := y2023.SolveD05P02(file)
		if err != nil {
			t.Fatal(err)
		}

		expected := "20358599"

		if got != expected {
			t.Fatalf("[expected %s] [actual %s]", expected, got)
		}
	})
}
