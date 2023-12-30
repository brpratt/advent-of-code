package y2023

import (
	"bufio"
	"io"
	"strconv"

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

// func printPlatform(g grid.Grid[rune]) {
// 	for y := 0; y < len(g); y++ {
// 		row := g.Row(y)
// 		for _, v := range row {
// 			fmt.Print(string(v))
// 		}
// 		fmt.Println()
// 	}
// }

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

func SolveD14P01(r io.Reader) (string, error) {
	var platform = parsePlatform(r)
	tilt(platform, grid.Up)
	return strconv.Itoa(northLoad(platform)), nil
}

func SolveD14P02(r io.Reader) (string, error) {
	var platform = parsePlatform(r)
	for n := 0; n < 1000000000; n++ {
		spin(platform)
	}
	tilt(platform, grid.Up)
	return strconv.Itoa(northLoad(platform)), nil
}
