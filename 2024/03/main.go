package main

import (
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/parse"
)

var (
	multPattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doPattern   = regexp.MustCompile(`do\(\)`)
	dontPattern = regexp.MustCompile(`don't\(\)`)
)

func part01(memory []byte) int {
	matches := multPattern.FindAllStringSubmatch(string(memory), -1)

	var sum int
	for _, match := range matches {
		a, b := match[1], match[2]
		sum += parse.MustAtoi(a) * parse.MustAtoi(b)
	}

	return sum
}

func indexOrDefault(matches [][]int, i int) int {
	if len(matches) == 0 {
		return i
	}

	return matches[0][0]
}

func part02(memory []byte) int {
	var (
		multMatches = multPattern.FindAllStringSubmatchIndex(string(memory), -1)
		doMatches   = doPattern.FindAllStringIndex(string(memory), -1)
		dontMatches = dontPattern.FindAllStringIndex(string(memory), -1)
		sum         = 0
		enabled     = true
	)

	for len(multMatches) > 0 {
		multIndex := indexOrDefault(multMatches, 0)
		doIndex := indexOrDefault(doMatches, multIndex+1)
		dontIndex := indexOrDefault(dontMatches, multIndex+1)

		switch {
		case multIndex < dontIndex && multIndex < doIndex:
			if !enabled {
				multMatches = multMatches[1:]
				continue
			}

			a, b := parse.MustAtoi(string(memory[multMatches[0][2]:multMatches[0][3]])), parse.MustAtoi(string(memory[multMatches[0][4]:multMatches[0][5]]))
			sum += a * b
			multMatches = multMatches[1:]
		case dontIndex < doIndex && dontIndex < multIndex:
			enabled = false
			dontMatches = dontMatches[1:]
		case doIndex < dontIndex && doIndex < multIndex:
			enabled = true
			doMatches = doMatches[1:]
		}
	}

	return sum
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	memory := file.Must(io.ReadAll(input))

	fmt.Println(part01(memory))
	fmt.Println(part02(memory))
}
