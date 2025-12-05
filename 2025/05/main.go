package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/parse"
)

type idrange struct {
	start int
	end   int
}

func (idr idrange) includes(id int) bool {
	return id >= idr.start && id <= idr.end
}

func (idr idrange) count() int {
	return (idr.end - idr.start) + 1
}

func overlaps(a, b idrange) bool {
	if b.end < a.start {
		return false
	}

	if b.start > a.end {
		return false
	}

	return true
}

func merge(a, b idrange) idrange {
	return idrange{
		start: min(b.start, a.start),
		end:   max(b.end, a.end),
	}
}

func parseRange(input string) idrange {
	var start, end int

	_, err := fmt.Sscanf(input, "%d-%d", &start, &end)
	if err != nil {
		panic("failed to parse ID range")
	}

	return idrange{start: start, end: end}
}

func isFresh(id int, idranges []idrange) bool {
	for _, idrange := range idranges {
		if idrange.includes(id) {
			return true
		}
	}

	return false
}

func parseInput(r io.Reader) ([]idrange, []int) {
	scanner := bufio.NewScanner(r)

	var idranges []idrange
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		idranges = append(idranges, parseRange(line))
	}

	var ids []int
	for scanner.Scan() {
		line := scanner.Text()
		ids = append(ids, parse.MustAtoi(line))
	}

	return idranges, ids
}

func part01(idranges []idrange, ids []int) int {
	var fresh int

	for _, id := range ids {
		if isFresh(id, idranges) {
			fresh++
		}
	}

	return fresh
}

func appendWithMerge(idranges []idrange, idr idrange) []idrange {
	for i := range idranges {
		if overlaps(idr, idranges[i]) {
			idranges[i] = merge(idr, idranges[i])
			return idranges
		}
	}

	idranges = append(idranges, idr)
	return idranges
}

func squash(idranges []idrange) []idrange {
	squashed := make([]idrange, len(idranges))
	copy(squashed, idranges)

	for {
		var merged []idrange
		for _, idrange := range squashed {
			merged = appendWithMerge(merged, idrange)
		}

		if len(squashed) == len(merged) {
			break
		}

		squashed = merged
	}

	return squashed
}

func part02(idranges []idrange) int {
	var fresh int

	squashed := squash(idranges)
	for _, idrange := range squashed {
		fresh += idrange.count()
	}

	return fresh
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	idranges, ids := parseInput(input)
	fmt.Println(part01(idranges, ids))
	fmt.Println(part02(idranges))
}
