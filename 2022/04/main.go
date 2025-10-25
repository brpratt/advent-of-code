package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/brpratt/advent-of-code/file"
)

type interval struct {
	min int
	max int
}

func contains(int1, int2 interval) bool {
	if int1.min <= int2.min && int1.max >= int2.max {
		return true
	}

	return int2.min <= int1.min && int2.max >= int1.max
}

func overlaps(int1, int2 interval) bool {
	return !(int1.max < int2.min || int2.max < int1.min)
}

func solve(input io.Reader) (string, string) {
	scanner := bufio.NewScanner(input)

	partA := 0
	partB := 0

	for scanner.Scan() {
		line := scanner.Text()

		var int1 interval
		var int2 interval

		fmt.Sscanf(line, "%d-%d,%d-%d", &int1.min, &int1.max, &int2.min, &int2.max)

		if contains(int1, int2) {
			partA += 1
		}

		if overlaps(int1, int2) {
			partB += 1
		}
	}

	return strconv.Itoa(partA), strconv.Itoa(partB)
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	part1, part2 := solve(input)
	fmt.Println(part1)
	fmt.Println(part2)
}
