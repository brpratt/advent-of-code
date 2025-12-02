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

func repeat(n int) int {
	return (n * pow(10, digits(n))) + n
}

func pivot(n int) int {
	ds := digits(n)

	if ds%2 == 0 {
		return n / pow(10, ds/2)
	}

	return pow(10, ds/2)
}

type idrange struct {
	start int
	stop  int
}

func (r *idrange) invalidIDs() iter.Seq[int] {
	p := pivot(r.start)
	for repeat(p) < r.start {
		p++
	}

	next := repeat(p)

	return func(yield func(int) bool) {
		for next <= r.stop {
			if !yield(next) {
				return
			}

			p++
			next = repeat(p)
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
		for id := range idrange.invalidIDs() {
			sum += id
		}
	}

	return sum
}

func main() {
	input := file.Must(file.ReadLines("input.txt"))[0]
	idranges := parseRanges(input)

	fmt.Println(part01(idranges))
}
