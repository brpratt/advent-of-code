package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brpratt/advent-of-code/file"
)

type cubeset struct {
	red   int
	green int
	blue  int
}

func (c cubeset) power() int {
	return c.red * c.green * c.blue
}

type game struct {
	id   int
	sets []cubeset
}

func (g game) minimum() (c cubeset) {
	for _, set := range g.sets {
		if set.red > c.red {
			c.red = set.red
		}
		if set.green > c.green {
			c.green = set.green
		}
		if set.blue > c.blue {
			c.blue = set.blue
		}
	}
	return
}

func main() {
	lines := file.Must(file.ReadLines("input.txt"))

	fmt.Println(part01(lines))
	fmt.Println(part02(lines))
}

func parseCubeset(s string) (c cubeset) {
	for _, cubestr := range strings.Split(s, ", ") {
		num, color, _ := strings.Cut(cubestr, " ")
		switch color {
		case "red":
			c.red, _ = strconv.Atoi(num)
		case "green":
			c.green, _ = strconv.Atoi(num)
		case "blue":
			c.blue, _ = strconv.Atoi(num)
		}
	}

	return
}

func parseGame(s string) (g game) {
	header, body, _ := strings.Cut(s, ": ")
	_, id, _ := strings.Cut(header, " ")
	g.id, _ = strconv.Atoi(id)

	for _, round := range strings.Split(body, "; ") {
		g.sets = append(g.sets, parseCubeset(round))
	}

	return
}

func possible(g game, c cubeset) bool {
	for _, set := range g.sets {
		if set.red > c.red || set.green > c.green || set.blue > c.blue {
			return false
		}
	}

	return true
}

func part01(lines []string) int {
	bound := cubeset{
		red:   12,
		green: 13,
		blue:  14,
	}

	sum := 0

	for _, line := range lines {
		game := parseGame(line)
		if possible(game, bound) {
			sum += game.id
		}
	}

	return sum
}

func part02(lines []string) int {
	sum := 0

	for _, line := range lines {
		game := parseGame(line)
		sum += game.minimum().power()
	}

	return sum
}
