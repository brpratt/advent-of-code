package main

import (
	"fmt"
	"iter"
	"strings"

	"github.com/brpratt/advent-of-code/file"
)

func pow(n, s int) int {
	if s <= 0 {
		return 1
	}

	m := n
	for s > 1 {
		m *= n
		s--
	}

	return m
}

func digits(n int) int {
	var count int

	for n > 0 {
		n /= 10
		count++
	}

	return count
}

func repeat(n, count int) int {
	var (
		ds = digits(n)
		m  = n
	)

	for count > 1 {
		m = m*pow(10, ds) + n
		count--
	}

	return m
}

type idrange struct {
	start int
	stop  int
}

func (r *idrange) invalidIDs(extended bool) iter.Seq[int] {
	pivot := 1

	return func(yield func(int) bool) {
		seen := make(map[int]bool)

		for repeat(pivot, 2) <= r.stop {
			repetition := 2

			for {
				next := repeat(pivot, repetition)
				if next > r.stop {
					break
				}

				if next >= r.start && !seen[next] {
					seen[next] = true
					if !yield(next) {
						return
					}
				}

				if !extended {
					break
				}

				repetition++
			}

			pivot++
		}
	}
}

func parseRange(input string) idrange {
	var start, stop int

	_, err := fmt.Sscanf(input, "%d-%d", &start, &stop)
	if err != nil {
		panic("failed to parse ID range")
	}

	return idrange{start, stop}
}

func parseRanges(input string) []idrange {
	sections := strings.Split(input, ",")
	ranges := make([]idrange, len(sections))

	for i, section := range sections {
		ranges[i] = parseRange(section)
	}

	return ranges
}

func part01(idranges []idrange) int {
	var sum int

	for _, idrange := range idranges {
		for id := range idrange.invalidIDs(false) {
			sum += id
		}
	}

	return sum
}

func part02(idranges []idrange) int {
	var sum int

	for _, idrange := range idranges {
		for id := range idrange.invalidIDs(true) {
			sum += id
		}
	}

	return sum
}

func main() {
	input := file.Must(file.ReadLines("input.txt"))[0]
	idranges := parseRanges(input)

	fmt.Println(part01(idranges))
	fmt.Println(part02(idranges))
}
