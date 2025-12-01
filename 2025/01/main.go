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

func part01(rotations []int) int {
	zeroCount := 0
	pos := 50
	for _, rotation := range rotations {
		pos += rotation
		pos %= 100
		if pos == 0 {
			zeroCount++
		}
	}
	return zeroCount
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	rotations := parseRotations(input)
	fmt.Println(part01(rotations))
}
