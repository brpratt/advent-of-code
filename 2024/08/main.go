package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/grid"
)

type antenna struct {
	pos grid.Point
	sym rune
}

type plot struct {
	dim      grid.Point
	antennas []antenna
}

func parseInput(r io.Reader) plot {
	var x int
	var y int
	var p plot

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		x = len(line)

		for i, c := range line {
			if c == '.' {
				continue
			}

			p.antennas = append(p.antennas, antenna{
				pos: grid.Point{X: i, Y: y},
				sym: c,
			})
		}
		y++
	}

	p.dim = grid.Point{X: x, Y: y}
	return p
}

func inBound(dim grid.Point, p grid.Point) bool {
	return p.X >= 0 && p.X < dim.X && p.Y >= 0 && p.Y < dim.Y
}

func part01(p plot) int {
	antinnodes := make(map[grid.Point]struct{})

	for i, a1 := range p.antennas {
		for _, a2 := range p.antennas[i+1:] {
			if a1.sym != a2.sym {
				continue
			}

			xdiff := a1.pos.X - a2.pos.X
			ydiff := a1.pos.Y - a2.pos.Y

			an1 := grid.Point{X: a1.pos.X + xdiff, Y: a1.pos.Y + ydiff}
			an2 := grid.Point{X: a2.pos.X - xdiff, Y: a2.pos.Y - ydiff}

			if inBound(p.dim, an1) {
				antinnodes[an1] = struct{}{}
			}

			if inBound(p.dim, an2) {
				antinnodes[an2] = struct{}{}
			}
		}
	}

	return len(antinnodes)
}

func part02(p plot) int {
	antinnodes := make(map[grid.Point]struct{})

	for i, a1 := range p.antennas {
		antinnodes[a1.pos] = struct{}{}

		for _, a2 := range p.antennas[i+1:] {
			if a1.sym != a2.sym {
				continue
			}

			xdiff := a1.pos.X - a2.pos.X
			ydiff := a1.pos.Y - a2.pos.Y

			an := grid.Point{X: a1.pos.X + xdiff, Y: a1.pos.Y + ydiff}
			for inBound(p.dim, an) {
				antinnodes[an] = struct{}{}
				an.X += xdiff
				an.Y += ydiff
			}

			an = grid.Point{X: a2.pos.X - xdiff, Y: a2.pos.Y - ydiff}
			for inBound(p.dim, an) {
				antinnodes[an] = struct{}{}
				an.X -= xdiff
				an.Y -= ydiff
			}
		}
	}

	return len(antinnodes)
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	equations := parseInput(input)

	fmt.Println(part01(equations))
	fmt.Println(part02(equations))
}
