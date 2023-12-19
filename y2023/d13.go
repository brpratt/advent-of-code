package y2023

import (
	"bufio"
	"io"
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

type symmertry int

const (
	horizontal symmertry = iota
	vertical
)

type mirror struct {
	symmetry symmertry
	index    int
	diff     int
}

func horizontalDiff(g grid.Grid[rune], y int) int {
	var count int

	if y < 0 || y >= len(g)-1 {
		return -1
	}

	for n := 0; y-n >= 0 && y+n+1 < len(g); n++ {
		row1 := g.Row(y - n)
		row2 := g.Row(y + n + 1)

		for x := range row1 {
			if row1[x] != row2[x] {
				count++
			}
		}
	}

	return count
}

func verticalDiff(g grid.Grid[rune], x int) int {
	var count int

	if x < 0 || x >= len(g[0])-1 {
		return -1
	}

	for n := 0; x-n >= 0 && x+n+1 < len(g[0]); n++ {
		column1 := g.Column(x - n)
		column2 := g.Column(x + n + 1)

		for y := range column1 {
			if column1[y] != column2[y] {
				count++
			}
		}
	}

	return count
}

func mirrors(g grid.Grid[rune]) []mirror {
	var mirrors []mirror

	for y := 0; y < len(g)-1; y++ {
		m := mirror{
			symmetry: horizontal,
			index:    y,
			diff:     horizontalDiff(g, y),
		}

		mirrors = append(mirrors, m)
	}

	for x := 0; x < len(g[0])-1; x++ {
		m := mirror{
			symmetry: vertical,
			index:    x,
			diff:     verticalDiff(g, x),
		}

		mirrors = append(mirrors, m)
	}

	return mirrors
}

func SolveD13P01(r io.Reader) (string, error) {
	walks := parseWalks(r)

	var count int
	for _, walk := range walks {
		for _, m := range mirrors(walk) {
			if m.diff != 0 {
				continue
			}

			switch m.symmetry {
			case horizontal:
				count += 100 * (m.index + 1)
			case vertical:
				count += m.index + 1
			}
		}
	}

	return strconv.Itoa(count), nil
}

func SolveD13P02(r io.Reader) (string, error) {
	walks := parseWalks(r)

	var count int
	for _, walk := range walks {
		for _, m := range mirrors(walk) {
			if m.diff != 1 {
				continue
			}

			switch m.symmetry {
			case horizontal:
				count += 100 * (m.index + 1)
			case vertical:
				count += m.index + 1
			}
		}
	}

	return strconv.Itoa(count), nil
}
