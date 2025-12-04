package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/grid"
)

func parseGrid(r io.Reader) grid.Grid[byte] {
	var g grid.Grid[byte]

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := []byte(scanner.Text())
		g = append(g, row)
	}

	return g
}

func adjacentPoints(p grid.Point) []grid.Point {
	return []grid.Point{
		grid.Move(p, grid.Up),
		grid.Move(p, grid.UpRight),
		grid.Move(p, grid.Right),
		grid.Move(p, grid.DownRight),
		grid.Move(p, grid.Down),
		grid.Move(p, grid.DownLeft),
		grid.Move(p, grid.Left),
		grid.Move(p, grid.UpLeft),
	}
}

func adjancentRolls(p grid.Point, g grid.Grid[byte]) int {
	var count int

	for _, ap := range adjacentPoints(p) {
		value, _ := g.Value(ap)

		if value == '@' {
			count++
		}
	}

	return count
}

func accessibleRolls(g grid.Grid[byte]) []grid.Point {
	accessible := make([]grid.Point, 0)

	for y := range g {
		for x := range len(g[y]) {
			p := grid.Point{X: x, Y: y}

			value, _ := g.Value(p)
			if value == '@' && adjancentRolls(p, g) < 4 {
				accessible = append(accessible, p)
			}
		}
	}

	return accessible
}

func part01(g grid.Grid[byte]) int {
	return len(accessibleRolls(g))
}

func part02(g grid.Grid[byte]) int {
	var count int
	ng := g.Copy()

	for {
		accessible := accessibleRolls(ng)

		if len(accessible) == 0 {
			break
		}

		count += len(accessible)

		for _, p := range accessible {
			ng.SetValue(p, '.')
		}
	}

	return count
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	g := parseGrid(input)

	fmt.Println(part01(g))
	fmt.Println(part02(g))
}
