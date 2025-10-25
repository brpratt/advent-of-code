package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/brpratt/advent-of-code/file"
)

var aPoints = map[string]int{
	"A X": 4,
	"B X": 1,
	"C X": 7,
	"A Y": 8,
	"B Y": 5,
	"C Y": 2,
	"A Z": 3,
	"B Z": 9,
	"C Z": 6,
}

var bPoints = map[string]int{
	"A X": 3,
	"B X": 1,
	"C X": 2,
	"A Y": 4,
	"B Y": 5,
	"C Y": 6,
	"A Z": 8,
	"B Z": 9,
	"C Z": 7,
}

func solve(input io.Reader) (string, string) {
	scanner := bufio.NewScanner(input)

	aSum := 0
	bSum := 0

	for scanner.Scan() {
		round := scanner.Text()

		aSum += aPoints[round]
		bSum += bPoints[round]
	}

	return strconv.Itoa(aSum), strconv.Itoa(bSum)
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	part1, part2 := solve(input)
	fmt.Println(part1)
	fmt.Println(part2)
}
