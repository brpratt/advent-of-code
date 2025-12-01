package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/parse"
)

func parseRotation(line string) int {
	number := parse.MustAtoi(line[1:])

	if line[0] == 'R' {
		return number
	}

	return -number
}

func parseRotations(r io.Reader) []int {
	rotations := make([]int, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		rotation := parseRotation(scanner.Text())
		rotations = append(rotations, rotation)
	}

	return rotations
}

func rotate(pos, rotation int) (int, int) {
	var (
		startedAtZero     = pos == 0
		passesThroughZero = 0
	)

	pos += rotation

	for pos >= 100 {
		pos -= 100
		passesThroughZero++
	}

	for pos < 0 {
		pos += 100
		passesThroughZero++
	}

	if rotation < 0 {
		// If we started at zero, any rotation left will cause us to wrap around and "see"
		// a zero in the preceding 'for' loop. Since we only want to count new zeros, we
		// need to remove this phantom zero.
		if startedAtZero {
			passesThroughZero--
		}

		// If we're now at zero and we rotated left, the preceding 'for' loop wouldn't see
		// this zero since position never became negative.
		if pos == 0 {
			passesThroughZero++
		}

	}

	return pos, passesThroughZero
}

func part01(rotations []int) int {
	var (
		count = 0
		pos   = 50
	)

	for _, rotation := range rotations {
		pos, _ = rotate(pos, rotation)
		if pos == 0 {
			count++
		}
	}

	return count
}

func part02(rotations []int) int {
	var (
		count             = 0
		pos               = 50
		passesThroughZero = 0
	)

	for _, rotation := range rotations {
		pos, passesThroughZero = rotate(pos, rotation)
		count += passesThroughZero
	}

	return count
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	rotations := parseRotations(input)
	fmt.Println(part01(rotations))
	fmt.Println(part02(rotations))
}
