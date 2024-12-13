package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/grid"
)

func parseInput(r io.Reader) grid.Grid[byte] {
	var g grid.Grid[byte]

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]byte, len(line))

		for i := range line {
			row[i] = line[i]
		}

		g = append(g, row)
	}

	return g
}

type plot struct {
	area      int
	perimeter int
}

func cardinal(p grid.Point) []grid.Point {
	return []grid.Point{
		grid.Move(p, grid.Up),
		grid.Move(p, grid.Right),
		grid.Move(p, grid.Down),
		grid.Move(p, grid.Left),
	}
}

func part01(g grid.Grid[byte]) int {
	var plots []plot
	var visited grid.Grid[bool] = make(grid.Grid[bool], 0)

	for y := range g {
		row := make([]bool, len(g[y]))
		visited = append(visited, row)
	}

	for y := range g {
		for x := range g[y] {
			if visited[y][x] {
				continue
			}

			var plot plot
			var stack []grid.Point = []grid.Point{{Y: y, X: x}}
			visited[y][x] = true

			for len(stack) > 0 {
				p := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				plot.area++

				for _, cp := range cardinal(p) {
					if !g.InBound(cp) {
						plot.perimeter++
						continue
					}

					if g[cp.Y][cp.X] != g[p.Y][p.X] {
						plot.perimeter++
						continue
					}

					if !visited[cp.Y][cp.X] {
						stack = append(stack, cp)
						visited[cp.Y][cp.X] = true
					}
				}
			}

			plots = append(plots, plot)
		}
	}

	var sum int
	for _, plot := range plots {
		sum += plot.area * plot.perimeter
	}
	return sum
}

func part02(g grid.Grid[byte]) int {
	return 0
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	garden := parseInput(input)

	fmt.Println(part01(garden))
	fmt.Println(part02(garden))
}
