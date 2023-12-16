package y2023

import (
	"bufio"
	"io"
	"slices"
	"strconv"
)

type direction int

const (
	up direction = iota
	right
	down
	left
)

// determine if a connects with b
// d indicates the direction from a to b
func connected(a, b rune, d direction) bool {
	allowedDirections := map[direction][]rune{
		up:    {'S', '|', 'L', 'J'},
		down:  {'S', '|', 'F', '7'},
		left:  {'S', '-', 'J', '7'},
		right: {'S', '-', 'F', 'L'},
	}

	if !slices.Contains(allowedDirections[d], a) {
		return false
	}

	allowedRunes := map[direction][]rune{
		up:    {'S', '|', '7', 'F'},
		down:  {'S', '|', 'J', 'L'},
		left:  {'S', '-', 'F', 'L'},
		right: {'S', '-', 'J', '7'},
	}

	return slices.Contains(allowedRunes[d], b)
}

func parseField(r io.Reader) (f [][]rune) {
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

type grid[T any] [][]T

func (g grid[any]) contains(p point) bool {
	return p.y >= 0 && p.x >= 0 && p.y < len(g) && p.x < len(g[p.y])
}

func (g grid[any]) move(p point, d direction) (newp point, ok bool) {
	switch d {
	case up:
		newp.x = p.x
		newp.y = p.y - 1
	case down:
		newp.x = p.x
		newp.y = p.y + 1
	case left:
		newp.x = p.x - 1
		newp.y = p.y
	case right:
		newp.x = p.x + 1
		newp.y = p.y
	}

	ok = g.contains(newp)
	return
}

func (g grid[any]) emptyMask() [][]bool {
	mask := make([][]bool, len(g))

	for y, row := range g {
		mask[y] = make([]bool, len(row))
	}

	return mask
}

func move(f [][]rune, p point, d direction) (point, bool) {
	newp, ok := grid[rune](f).move(p, d)
	if !ok {
		return newp, false
	}

	tileA := f[p.y][p.x]
	tileB := f[newp.y][newp.x]

	return newp, connected(tileA, tileB, d)
}

type step struct {
	p point
	d direction
}

func walk(f [][]rune, start point) chan step {
	steps := make(chan step)

	go func() {
		curr := start
		visited := grid[rune](f).emptyMask()

		visited[curr.y][curr.x] = true

		var end bool
		for !end {
			end = true
			for d := up; d <= left; d++ {
				if next, ok := move(f, curr, d); ok && !visited[next.y][next.x] {
					end = false
					curr = next
					visited[curr.y][curr.x] = true
					steps <- step{curr, d}
					break
				}
			}
		}

		close(steps)
	}()

	return steps
}

func findStart(f [][]rune) (p point) {
	for y := 0; y < len(f); y++ {
		for x := 0; x < len(f[y]); x++ {
			if f[y][x] == 'S' {
				p.y = y
				p.x = x
				return
			}
		}
	}

	return
}

func mainLoop(f [][]rune) [][]bool {
	loop := grid[rune](f).emptyMask()

	start := findStart(f)
	loop[start.y][start.x] = true

	steps := walk(f, start)
	for {
		if step, ok := <-steps; ok {
			loop[step.p.y][step.p.x] = true
		} else {
			break
		}
	}

	return loop
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

func flood(points [][]bool, boundary [][]bool) {
	check := make([]point, 0, 10)
	for y := 0; y < len(points); y++ {
		for x := 0; x < len(points[y]); x++ {
			if points[y][x] {
				check = append(check, point{y: y, x: x})
			}
		}
	}

	for len(check) != 0 {
		cp := check[0]
		check = check[1:]

		points[cp.y][cp.x] = true

		// up
		if cp.y-1 >= 0 && !points[cp.y-1][cp.x] && !boundary[cp.y-1][cp.x] {
			check = append(check, point{y: cp.y - 1, x: cp.x})
		}
		// right
		if cp.x+1 < len(points[cp.y]) && !points[cp.y][cp.x+1] && !boundary[cp.y][cp.x+1] {
			check = append(check, point{y: cp.y, x: cp.x + 1})
		}
		// down
		if cp.y+1 < len(points) && !points[cp.y+1][cp.x] && !boundary[cp.y+1][cp.x] {
			check = append(check, point{y: cp.y + 1, x: cp.x})
		}
		// left
		if cp.x-1 >= 0 && !points[cp.y][cp.x-1] && !boundary[cp.y][cp.x-1] {
			check = append(check, point{y: cp.y, x: cp.x - 1})
		}
	}
}

func SolveD10P02(r io.Reader) (string, error) {
	f := parseField(r)
	loop := mainLoop(f)

	var start point

findStart:
	for y := 0; y < len(loop); y++ {
		for x := 0; x < len(loop[y]); x++ {
			if loop[y][x] {
				start.y = y
				start.x = x
				break findStart
			}
		}
	}

	first := true
	prevd := up
	prevp := start
	steps := walk(f, start)
	inner := grid[rune](f).emptyMask()

	for {
		var step step
		var ok bool

		if step, ok = <-steps; !ok {
			break
		}

		if first {
			prevd = step.d
			first = false
		}

		var innerPoint point
		switch step.d {
		case up:
			innerPoint.y = prevp.y
			innerPoint.x = prevp.x + 1
		case right:
			innerPoint.y = prevp.y + 1
			innerPoint.x = prevp.x
		case down:
			innerPoint.y = prevp.y
			innerPoint.x = prevp.x - 1
		case left:
			innerPoint.y = prevp.y - 1
			innerPoint.x = prevp.x
		}

		if grid[bool](inner).contains(innerPoint) && !loop[innerPoint.y][innerPoint.x] {
			inner[innerPoint.y][innerPoint.x] = true
		}

		if step.d == right && prevd == down {
			innerPoint.y = prevp.y
			innerPoint.x = prevp.x - 1
			if grid[bool](inner).contains(innerPoint) && !loop[innerPoint.y][innerPoint.x] {
				inner[innerPoint.y][innerPoint.x] = true
			}
		}

		prevp = step.p
		prevd = step.d
	}

	flood(inner, loop)
	// printField(f, loop, inner)

	var count int
	for y := 0; y < len(inner); y++ {
		for x := 0; x < len(inner[y]); x++ {
			if inner[y][x] {
				count++
			}
		}
	}

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
