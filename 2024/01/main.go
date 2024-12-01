package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/parse"
)

func parseLists(r io.Reader) (left []int, right []int) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		left = append(left, parse.MustAtoi(fields[0]))
		right = append(right, parse.MustAtoi(fields[1]))
	}

	return
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func part01(left, right []int) int {
	leftSorted := slices.Sorted(slices.Values(left))
	rightSorted := slices.Sorted(slices.Values(right))

	var distance int

	for i := 0; i < len(leftSorted); i++ {
		distance += abs(leftSorted[i] - rightSorted[i])
	}

	return distance
}

func part02(left, right []int) int {
	appearance := make(map[int]int)
	for _, n := range right {
		appearance[n]++
	}

	var similarity int
	for _, n := range left {
		similarity += n * appearance[n]
	}

	return similarity
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	left, right := parseLists(input)
	fmt.Println(part01(left, right))
	fmt.Println(part02(left, right))
}
