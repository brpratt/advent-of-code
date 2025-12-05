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

func parseIDRange(input string) idrange {
	var start, stop int

	_, err := fmt.Sscanf(input, "%d-%d", &start, &stop)
	if err != nil {
		panic("failed to parse ID range")
	}

	return idrange{start, stop}
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

		idranges = append(idranges, parseIDRange(line))
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

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	idranges, ids := parseInput(input)
	fmt.Println(part01(idranges, ids))
}
