package main

import (
	"fmt"
	"strings"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/grid"
	"github.com/brpratt/advent-of-code/parse"
)

type block []string

func (b block) width() int {
	var maxlen int

	for _, entry := range b {
		maxlen = max(maxlen, len(entry))
	}

	return maxlen
}

func (b block) height() int {
	return len(b)
}

func parseInput(lines []string) ([]block, []byte) {
	var (
		operators = make([]byte, 0)
		opline    = lines[len(lines)-1]
		opidx     = 0
		widths    = make([]int, 0)
	)

	for opidx < len(opline) {
		operators = append(operators, opline[opidx])

		nextopidx := opidx + 1
		for nextopidx < len(opline) && opline[nextopidx] == ' ' {
			nextopidx++
		}

		if nextopidx == len(opline) {
			nextopidx++
		}

		widths = append(widths, nextopidx-opidx-1)
		opidx = nextopidx
	}

	blocks := make([]block, len(widths))

	for _, line := range lines[:len(lines)-1] {
		offset := 0
		for nblock, width := range widths {
			blocks[nblock] = append(blocks[nblock], line[offset:offset+width])
			offset += width + 1
		}
	}

	return blocks, operators
}

func add(b block) int {
	var sum int

	for _, entry := range b {
		sum += parse.MustAtoi(strings.TrimSpace(entry))
	}

	return sum
}

func multiply(b block) int {
	var product int = 1

	for _, entry := range b {
		product *= parse.MustAtoi(strings.TrimSpace(entry))
	}

	return product
}

func compute(b block, op byte) int {
	if op == '+' {
		return add(b)
	}

	return multiply(b)
}

func rotate(b block) block {
	var (
		rheight = b.width()
		rwidth  = b.height()
		g       = make(grid.Grid[byte], rheight)
	)

	for i := range rheight {
		g[i] = make([]byte, rwidth)
	}

	for x, entry := range b {
		for y, ch := range []byte(entry) {
			g[rheight-y-1][x] = ch
		}
	}

	rotated := make(block, rheight)
	for i := range rheight {
		rotated[i] = string(g[i])
	}

	return rotated
}

func part01(blocks []block, operators []byte) int {
	var sum int

	for i := range operators {
		sum += compute(blocks[i], operators[i])
	}

	return sum
}

func part02(blocks []block, operators []byte) int {
	var sum int

	for i := range operators {
		sum += compute(rotate(blocks[i]), operators[i])
	}

	return sum
}

func main() {
	input := file.Must(file.ReadLines("input.txt"))
	blocks, operators := parseInput(input)

	fmt.Println(part01(blocks, operators))
	fmt.Println(part02(blocks, operators))
}
