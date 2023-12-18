package y2023

import (
	"bufio"
	"io"
	"slices"
	"strconv"

	"github.com/brpratt/advent-of-code/grid"
)

type universe struct {
	image    grid.Grid[rune]
	galaxies []grid.Point
	emptyYs  []int
	emptyXs  []int
}

func parseUniverse(r io.Reader) (u universe) {
	u.image = make([][]rune, 0)

	var galaxyInX []bool

	y := 0
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		row := []rune(scanner.Text())

		if galaxyInX == nil {
			galaxyInX = make([]bool, len(row))
		}

		var galaxyInY bool

		for x, tile := range row {
			if tile != '#' {
				continue
			}

			galaxyInY = true
			galaxyInX[x] = true
			u.galaxies = append(u.galaxies, grid.Point{Y: y, X: x})
		}

		if !galaxyInY {
			u.emptyYs = append(u.emptyYs, y)
		}

		y++
	}

	for x := range galaxyInX {
		if !galaxyInX[x] {
			u.emptyXs = append(u.emptyXs, x)
		}
	}

	return
}

func distance(u universe, a, b grid.Point, scale int) int {
	if a.Y < b.Y {
		a.Y, b.Y = b.Y, a.Y
	}

	offsetY := a.Y - b.Y

	if a.X < b.X {
		a.X, b.X = b.X, a.X
	}

	offsetX := a.X - b.X

	for y := b.Y + 1; y < a.Y; y++ {
		if slices.Contains(u.emptyYs, y) {
			offsetY += (scale - 1)
		}
	}

	for x := b.X + 1; x < a.X; x++ {
		if slices.Contains(u.emptyXs, x) {
			offsetX += (scale - 1)
		}
	}

	return offsetY + offsetX
}

func SolveD11P01(r io.Reader) (string, error) {
	universe := parseUniverse(r)

	var distances []int

	for from := range universe.galaxies {
		for to := from + 1; to < len(universe.galaxies); to++ {
			fromGalaxy := universe.galaxies[from]
			toGalaxy := universe.galaxies[to]
			d := distance(universe, fromGalaxy, toGalaxy, 2)
			distances = append(distances, d)
		}
	}

	var sum int
	for _, d := range distances {
		sum += d
	}

	return strconv.Itoa(sum), nil
}

func SolveD11P02(r io.Reader) (string, error) {
	universe := parseUniverse(r)

	var distances []int

	for from := range universe.galaxies {
		for to := from + 1; to < len(universe.galaxies); to++ {
			fromGalaxy := universe.galaxies[from]
			toGalaxy := universe.galaxies[to]
			d := distance(universe, fromGalaxy, toGalaxy, 1000000)
			distances = append(distances, d)
		}
	}

	var sum int
	for _, d := range distances {
		sum += d
	}

	return strconv.Itoa(sum), nil
}
