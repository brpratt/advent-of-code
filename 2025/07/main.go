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

func start(g grid.Grid[byte]) grid.Point {
	for x, ch := range g[0] {
		if ch == 'S' {
			return grid.Point{X: x, Y: 0}
		}
	}

	panic("no starting point")
}

func part01(g grid.Grid[byte]) int {
	gcopy := g.Copy()
	stack := []grid.Point{start(gcopy)}
	count := 0

	for len(stack) != 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		down := grid.Move(point, grid.Down)
		value, inBound := gcopy.Value(down)
		if !inBound || value == '|' {
			continue
		}

		if value == '.' {
			gcopy.SetValue(down, '|')
			stack = append(stack, down)
			continue
		}

		// we hit a splitter
		count++

		downleft := grid.Move(point, grid.DownLeft)
		if gcopy.InBound(downleft) {
			stack = append(stack, downleft)
		}

		downright := grid.Move(point, grid.DownRight)
		if gcopy.InBound(downright) {
			stack = append(stack, downright)
		}
	}

	return count
}

func walk(g grid.Grid[byte], p grid.Point, memo map[grid.Point]int) int {
	if count, ok := memo[p]; ok {
		return count
	}

	if !g.InBound(p) {
		memo[p] = 0
		return 0
	}

	next := grid.Move(p, grid.Down)
	value, inBound := g.Value(next)

	for inBound && value == '.' {
		next = grid.Move(next, grid.Down)
		value, inBound = g.Value(next)
	}

	if !inBound {
		memo[p] = 1
		return 1
	}

	// we hit a splitter

	left := grid.Move(next, grid.Left)
	right := grid.Move(next, grid.Right)

	count := walk(g, left, memo) + walk(g, right, memo)
	memo[p] = count
	return count
}

func part02(g grid.Grid[byte]) int {
	return walk(g, start(g), make(map[grid.Point]int))
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	grid := parseGrid(input)
	fmt.Println(part01(grid))
	fmt.Println(part02(grid))
}
