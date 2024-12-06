package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/grid"
)

func parseInput(r io.Reader) (grid.Point, grid.Grid[byte]) {
	var g grid.Grid[byte] = make(grid.Grid[byte], 0)
	var start grid.Point
	var y int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		row := []byte(line)

		if x := slices.Index(row, '^'); x != -1 {
			start = grid.Point{X: x, Y: y}
			row[x] = '.'
		}

		g = append(g, row)
		y++
	}

	return start, g
}

func turnBack(d grid.Direction) grid.Direction {
	switch d {
	case grid.Up:
		return grid.Down
	case grid.Right:
		return grid.Left
	case grid.Down:
		return grid.Up
	case grid.Left:
		return grid.Right
	}

	return d
}

func turnRight(d grid.Direction) grid.Direction {
	switch d {
	case grid.Up:
		return grid.Right
	case grid.Right:
		return grid.Down
	case grid.Down:
		return grid.Left
	case grid.Left:
		return grid.Up
	}

	return d
}

func part01(p grid.Point, g grid.Grid[byte]) int {
	var d grid.Direction = grid.Up
	var visited = make(grid.Grid[bool], 0)
	var visitedCount int

	for row := range g {
		visited = append(visited, make([]bool, len(g[row])))
	}

	for g.InBound(p) {
		if g[p.Y][p.X] == '#' {
			d = turnBack(d)
			p = grid.Move(p, d)
			d = turnBack(d)
			d = turnRight(d)
			p = grid.Move(p, d)
			continue
		}

		if !visited[p.Y][p.X] {
			visited[p.Y][p.X] = true
			visitedCount++
		}

		p = grid.Move(p, d)
	}

	return visitedCount
}

func isLoop(p grid.Point, g grid.Grid[byte]) bool {
	var d grid.Direction = grid.Up
	var visited = make([][]int, 0)

	for row := range g {
		visited = append(visited, make([]int, len(g[row])))
	}

	for g.InBound(p) {
		if g[p.Y][p.X] == '#' {
			d = turnBack(d)
			p = grid.Move(p, d)
			d = turnBack(d)
			d = turnRight(d)
			p = grid.Move(p, d)
			continue
		}

		if visited[p.Y][p.X]&(1<<d) != 0 {
			return true
		}

		visited[p.Y][p.X] |= 1 << d
		p = grid.Move(p, d)
	}

	return false
}

func part02(p grid.Point, g grid.Grid[byte]) int {
	var count int

	for y := range g {
		for x := range g[y] {
			if p.X == x && p.Y == y {
				continue
			}

			if g[y][x] == '#' {
				continue
			}

			g[y][x] = '#'
			if isLoop(p, g) {
				count++
			}
			g[y][x] = '.'
		}
	}

	return count
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	pageOrders, updates := parseInput(input)

	fmt.Println(part01(pageOrders, updates))
	fmt.Println(part02(pageOrders, updates))
}
