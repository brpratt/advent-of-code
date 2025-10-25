package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/brpratt/advent-of-code/file"
)

func value(r rune) int {
	if r >= 97 {
		return int(r) - 96
	}

	return int(r) - 38
}

type partA struct {
	sum int
}

func (p *partA) Add(line string) {
	sack1 := line[:len(line)/2]
	sack2 := line[len(line)/2:]

	items := map[rune]bool{}

	for _, item := range sack1 {
		items[item] = true
	}

	for _, item := range sack2 {
		if items[item] {
			p.sum += value(item)
			break
		}
	}
}

func (p *partA) Answer() string {
	return strconv.Itoa(p.sum)
}

type partB struct {
	sum   int
	count int
	sacks [3]string
}

func (p *partB) Add(line string) {
	p.sacks[p.count] = line
	p.count += 1

	if p.count%3 == 0 {
		p.count = 0
		uniq1 := map[rune]bool{}
		uniq2 := map[rune]bool{}

		for _, item := range p.sacks[0] {
			uniq1[item] = true
		}

		for _, item := range p.sacks[1] {
			if uniq1[item] {
				uniq2[item] = true
			}
		}

		for _, item := range p.sacks[2] {
			if uniq2[item] {
				p.sum += value(item)
				break
			}
		}
	}
}

func (p *partB) Answer() string {
	return strconv.Itoa(p.sum)
}

func solve(input io.Reader) (string, string) {
	scanner := bufio.NewScanner(input)

	partA := partA{}
	partB := partB{}

	for scanner.Scan() {
		line := scanner.Text()

		partA.Add(line)
		partB.Add(line)
	}

	return partA.Answer(), partB.Answer()
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	part1, part2 := solve(input)
	fmt.Println(part1)
	fmt.Println(part2)
}
