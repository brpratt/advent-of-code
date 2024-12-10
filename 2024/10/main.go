package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/grid"
)

func parseInput(r io.Reader) grid.Grid[int] {
	var g grid.Grid[int]

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, c := range line {
			row[i] = int(c - '0')
		}
		g = append(g, row)
	}

	return g
}

func cardinal(p grid.Point) []grid.Point {
	return []grid.Point{
		grid.Move(p, grid.Up),
		grid.Move(p, grid.Right),
		grid.Move(p, grid.Down),
		grid.Move(p, grid.Left),
	}
}

func score(p grid.Point, g grid.Grid[int]) int {
	seen := make(map[grid.Point]struct{})
	path := []grid.Point{p}

	for len(path) > 0 {
		var (
			top    = len(path) - 1
			p      = path[top]
			height = g[p.Y][p.X]
		)

		path = path[:top]

		if height == 9 {
			seen[p] = struct{}{}
			continue
		}

		for _, next := range cardinal(p) {
			nextHeight, ok := g.Value(next)
			if ok && nextHeight == height+1 {
				path = append(path, next)
			}
		}
	}

	return len(seen)
}

func part01(g grid.Grid[int]) int {
	sum := 0
	trailheads := g.FindAll(0)

	for _, th := range trailheads {
		sum += score(th, g)
	}

	return sum
}

func rating(p grid.Point, g grid.Grid[int]) int {
	count := 0
	path := []grid.Point{p}

	for len(path) > 0 {
		var (
			top    = len(path) - 1
			p      = path[top]
			height = g[p.Y][p.X]
		)

		path = path[:top]

		if height == 9 {
			count++
			continue
		}

		for _, next := range cardinal(p) {
			nextHeight, ok := g.Value(next)
			if ok && nextHeight == height+1 {
				path = append(path, next)
			}
		}
	}

	return count
}

func part02(g grid.Grid[int]) int {
	sum := 0
	trailheads := g.FindAll(0)

	for _, th := range trailheads {
		sum += rating(th, g)
	}

	return sum
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	fs := parseInput(input)

	fmt.Println(part01(fs))
	fmt.Println(part02(fs))
}
