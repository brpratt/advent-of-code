package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/parse"
)

func parseInput(r io.Reader) []int {
	stones := make([]int, 0)

	bytes, _ := io.ReadAll(r)
	for _, field := range strings.Fields(string(bytes)) {
		stones = append(stones, parse.MustAtoi(field))
	}

	return stones
}

func strlen(n int) int {
	str := strconv.Itoa(n)
	return len(str)
}

type node struct {
	stone int
	step  int
}

var cache = make(map[node]int)

func count(stone, step int) int {
	var total int
	n := node{stone: stone, step: step}

	if total, ok := cache[n]; ok {
		return total
	}

	switch {
	case step == 0:
		total = 1
	case stone == 0:
		total = count(1, step-1)
	case strlen(stone)%2 == 0:
		str := strconv.Itoa(stone)
		left := parse.MustAtoi(str[:len(str)/2])
		right := parse.MustAtoi(str[len(str)/2:])
		total = count(left, step-1) + count(right, step-1)
	default:
		total = count(stone*2024, step-1)
	}

	cache[node{stone: stone, step: step}] = total
	return total
}

func part01(stones []int) int {
	var sum int

	for _, stone := range stones {
		sum += count(stone, 25)
	}

	return sum
}

func part02(stones []int) int {
	var sum int

	for _, stone := range stones {
		sum += count(stone, 75)
	}

	return sum
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	stones := parseInput(input)

	fmt.Println(part01(stones))
	fmt.Println(part02(stones))
}
