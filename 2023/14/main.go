package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/grid"
)

func parsePlatform(r io.Reader) grid.Grid[rune] {
	var platform grid.Grid[rune]

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		platform = append(platform, row)
	}

	return platform
}

func hash(platform grid.Grid[rune]) string {
	var builder strings.Builder

	for y := 0; y < len(platform); y++ {
		builder.WriteString(string(platform[y]))
	}

	return builder.String()
}

func hydrate(platform grid.Grid[rune], h string) {
	var r = []rune(h)
	var ri int

	for y := 0; y < len(platform); y++ {
		for x := 0; x < len(platform[y]); x++ {
			platform[y][x] = r[ri]
			ri++
		}
	}
}

func tilt(g grid.Grid[rune], d grid.Direction) {
	var pin int
	var rocks int

	switch d {
	case grid.Up:
		for x := 0; x < len(g[0]); x++ {
			col := g.Column(x)
			for y, tile := range col {
				switch tile {
				case 'O':
					g[y][x] = '.'
					g[pin+rocks][x] = 'O'
					rocks++
				case '#':
					pin = y + 1
					rocks = 0
				}
			}
			pin = 0
			rocks = 0
		}
	case grid.Down:
		for x := 0; x < len(g[0]); x++ {
			col := g.Column(x)
			pin = len(col) - 1
			for y := len(col) - 1; y >= 0; y-- {
				tile := col[y]
				switch tile {
				case 'O':
					g[y][x] = '.'
					g[pin-rocks][x] = 'O'
					rocks++
				case '#':
					pin = y - 1
					rocks = 0
				}
			}
			pin = 0
			rocks = 0
		}
	case grid.Left:
		for y := 0; y < len(g); y++ {
			row := g.Row(y)
			for x, tile := range row {
				switch tile {
				case 'O':
					g[y][x] = '.'
					g[y][pin+rocks] = 'O'
					rocks++
				case '#':
					pin = x + 1
					rocks = 0
				}
			}
			pin = 0
			rocks = 0
		}
	case grid.Right:
		for y := 0; y < len(g); y++ {
			row := g.Row(y)
			pin = len(row) - 1
			for x := len(row) - 1; x >= 0; x-- {
				tile := row[x]
				switch tile {
				case 'O':
					g[y][x] = '.'
					g[y][pin-rocks] = 'O'
					rocks++
				case '#':
					pin = x - 1
					rocks = 0
				}
			}
			pin = 0
			rocks = 0
		}
	}
}

func spin(g grid.Grid[rune]) {
	tilt(g, grid.Up)
	tilt(g, grid.Left)
	tilt(g, grid.Down)
	tilt(g, grid.Right)
}

func spinUntilRepeat(platform grid.Grid[rune]) map[string]int {
	var count int

	patterns := make(map[string]int)
	prev := hash(platform)

	for {
		spin(platform)
		h := hash(platform)

		if _, ok := patterns[h]; ok {
			hydrate(platform, prev)
			return patterns
		}

		prev = h
		count++
		patterns[h] = count
	}

}

func northLoad(platform grid.Grid[rune]) int {
	var total int

	for x := 0; x < len(platform[0]); x++ {
		col := platform.Column(x)

		for y := 0; y < len(col); y++ {
			if col[y] == 'O' {
				total += len(col) - y
			}
		}
	}

	return total
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	platform := parsePlatform(input)
	fmt.Println(part01(platform))
	fmt.Println(part02(platform))
}

func part01(platform grid.Grid[rune]) int {
	tilt(platform, grid.Up)
	return northLoad(platform)
}

func part02(platform grid.Grid[rune]) int {
	initial := spinUntilRepeat(platform)
	cycle := spinUntilRepeat(platform)

	offset := (1000000000 - len(initial)) % len(cycle)
	for h, c := range cycle {
		if c == offset {
			hydrate(platform, h)
			break
		}
	}

	return northLoad(platform)
}
