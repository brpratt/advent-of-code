package y2023

import (
	"bufio"
	"io"
	"slices"
	"strconv"

	"github.com/brpratt/advent-of-code/grid"
)

// determine if a connects with b
// d indicates the direction from a to b
func connected(a, b rune, d grid.Direction) bool {
	allowedDirections := map[grid.Direction][]rune{
		grid.Up:    {'S', '|', 'L', 'J'},
		grid.Down:  {'S', '|', 'F', '7'},
		grid.Left:  {'S', '-', 'J', '7'},
		grid.Right: {'S', '-', 'F', 'L'},
	}

	if !slices.Contains(allowedDirections[d], a) {
		return false
	}

	allowedRunes := map[grid.Direction][]rune{
		grid.Up:    {'S', '|', '7', 'F'},
		grid.Down:  {'S', '|', 'J', 'L'},
		grid.Left:  {'S', '-', 'F', 'L'},
		grid.Right: {'S', '-', 'J', '7'},
	}

	return slices.Contains(allowedRunes[d], b)
}

func parseField(r io.Reader) (f grid.Grid[rune]) {
	f = make([][]rune, 0)

	y := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		f = append(f, row)
		y++
	}

	return
}

func mainLoop(f grid.Grid[rune]) grid.Grid[bool] {
	var loop grid.Grid[bool] = make([][]bool, len(f))
	for y := 0; y < len(f); y++ {
		loop[y] = make([]bool, len(f[y]))
	}

	curr, _ := f.Find('S')
	loop.SetValue(curr, true)

	var end bool
	for !end {
		end = true

		for d := grid.Up; d <= grid.Left; d++ {
			newp := grid.Move(curr, d)
			prev, _ := f.Value(curr)
			next, inBound := f.Value(newp)
			visited, _ := loop.Value(newp)

			if inBound && connected(prev, next, d) && !visited {
				end = false
				curr = newp
				loop.SetValue(curr, true)
			}
		}
	}

	return loop
}

func crossings(f grid.Grid[rune], loop grid.Grid[bool], p grid.Point) int {
	var count int
	var prevcorner rune

	if partOfLoop, _ := loop.Value(p); partOfLoop {
		return 0
	}

	newp := grid.Move(p, grid.Right)
	onLoop, inBound := loop.Value(newp)
	for inBound {
		if onLoop {
			tile, _ := f.Value(newp)
			switch {
			case tile == '|':
				count++
			case tile == 'F':
				prevcorner = tile
			case tile == '7':
				if prevcorner != 'F' {
					count++
				}
				prevcorner = tile
			case tile == 'J':
				if prevcorner != 'L' {
					count++
				}
				prevcorner = tile
			case tile == 'L':
				prevcorner = tile
			}
		}

		p = newp
		newp = grid.Move(p, grid.Right)
		onLoop, inBound = loop.Value(newp)
	}

	return count
}

func SolveD10P01(r io.Reader) (string, error) {
	f := parseField(r)
	loop := mainLoop(f)

	var count int
	for _, row := range loop {
		for _, yes := range row {
			if yes {
				count++
			}
		}
	}

	return strconv.Itoa((count + 1) / 2), nil
}

func SolveD10P02(r io.Reader) (string, error) {
	f := parseField(r)
	loop := mainLoop(f)

	sreplace := [4][4]rune{
		grid.Up: {
			grid.Right: 'L',
			grid.Down:  '|',
			grid.Left:  'J',
		},
		grid.Right: {
			grid.Down: 'F',
			grid.Left: '-',
		},
		grid.Down: {
			grid.Left: '7',
		},
	}

	var d1 grid.Direction
	var d2 grid.Direction
	var foundD1 bool

	spoint, _ := grid.Grid[rune](f).Find('S')
	for d := grid.Up; d <= grid.Left; d++ {
		newp := grid.Move(spoint, d)
		if partOfLoop, _ := loop.Value(newp); !partOfLoop {
			continue
		}
		if foundD1 {
			d2 = d
		} else {
			d1 = d
			foundD1 = true
		}
	}

	f.SetValue(spoint, sreplace[d1][d2])

	var inside grid.Grid[bool] = make([][]bool, len(f))
	for y := 0; y < len(f); y++ {
		inside[y] = make([]bool, len(f[y]))
	}

	var count int
	for y := range loop {
		for x := range loop[y] {
			p := grid.Point{Y: y, X: x}
			c := crossings(f, loop, p)
			if c%2 != 0 {
				inside.SetValue(p, true)
				count++
			}
		}
	}

	// printField(f, loop, inside)

	return strconv.Itoa(count), nil
}

/*
func printField(f [][]rune, loop [][]bool, inner [][]bool) {
	for y := range f {
		for x := range f[y] {
			switch {
			case loop[y][x]:
				switch f[y][x] {
				case 'F':
					fmt.Print("┌")
				case '7':
					fmt.Print("┐")
				case 'J':
					fmt.Print("┘")
				case 'L':
					fmt.Print("└")
				case '-':
					fmt.Print("─")
				case '|':
					fmt.Print("│")
				case 'S':
					fmt.Print("S")
				}
			case inner[y][x]:
				fmt.Print("X")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
*/
