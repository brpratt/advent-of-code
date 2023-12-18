package y2023

import (
	"bufio"
	"errors"
	"io"
	"slices"
	"strconv"

	"github.com/brpratt/advent-of-code/grid"
)

func parseWalks(r io.Reader) []grid.Grid[rune] {
	var walkn int
	var walks []grid.Grid[rune]

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := []rune(scanner.Text())

		if len(row) == 0 {
			walkn++
			continue
		}

		if len(walks) < walkn+1 {
			walks = append(walks, make(grid.Grid[rune], 0))
		}

		walks[walkn] = append(walks[walkn], row)
	}

	return walks
}

const (
	none = iota
	horizontal
	vertical
)

type mirror struct {
	symmetry int
	index    int
}

func findMirror(g grid.Grid[rune]) (m mirror) {
	// try finding horizontal mirrors
	for y := 0; y < len(g)-1; y++ {
		var found bool = true

		for n := 0; y-n >= 0 && y+n+1 < len(g); n++ {
			row1 := g.Row(y - n)
			row2 := g.Row(y + n + 1)
			if !slices.Equal(row1, row2) {
				found = false
				break
			}
		}

		if found {
			m.symmetry = horizontal
			m.index = y
			return
		}
	}

	// try finding vertical mirrors
	for x := 0; x < len(g[0])-1; x++ {
		var found bool = true

		for n := 0; x-n >= 0 && x+n+1 < len(g[0]); n++ {
			column1 := g.Column(x - n)
			column2 := g.Column(x + n + 1)
			if !slices.Equal(column1, column2) {
				found = false
				break
			}
		}

		if found {
			m.symmetry = vertical
			m.index = x
			return
		}
	}

	return
}

func SolveD13P01(r io.Reader) (string, error) {
	walks := parseWalks(r)

	var count int
	for _, walk := range walks {
		mirror := findMirror(walk)
		switch mirror.symmetry {
		case horizontal:
			count += 100 * (mirror.index + 1)
		case vertical:
			count += mirror.index + 1
		}
	}

	return strconv.Itoa(count), nil
}

func SolveD13P02(r io.Reader) (string, error) {
	return "", errors.New("not implemented")
}
