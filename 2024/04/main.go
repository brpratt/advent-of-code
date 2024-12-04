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

func wordFrom(g grid.Grid[byte], p grid.Point, d grid.Direction, n int) string {
	var builder = make([]byte, 0, len(g))

	for steps := 0; steps < n; steps++ {
		if letter, ok := g.Value(p); ok {
			builder = append(builder, letter)
			p = grid.Move(p, d)
		} else {
			break
		}
	}

	return string(builder)
}

func part01(g grid.Grid[byte]) int {
	var count int

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			if g[y][x] != 'X' {
				continue
			}

			for d := grid.Up; d <= grid.DownRight; d++ {
				if wordFrom(g, grid.Point{Y: y, X: x}, d, 4) == "XMAS" {
					count++
				}
			}
		}
	}

	return count
}

func isCrossMas(g grid.Grid[byte], p grid.Point) bool {
	for d := grid.Up; d <= grid.DownRight; d++ {
		if wordFrom(g, p, d, 4) == "XMAS" {
			return true
		}
	}

	return false
}

func part02(g grid.Grid[byte]) int {
	var count int

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			if g[y][x] != 'A' {
				continue
			}

			word := wordFrom(g, grid.Point{Y: y - 1, X: x - 1}, grid.DownRight, 3)
			if word != "MAS" && word != "SAM" {
				continue
			}

			word = wordFrom(g, grid.Point{Y: y + 1, X: x - 1}, grid.UpRight, 3)
			if word != "MAS" && word != "SAM" {
				continue
			}

			count++
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
