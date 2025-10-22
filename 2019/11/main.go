package main

import (
	"fmt"
	"strings"

	"github.com/brpratt/advent-of-code/2019/intcode"
	"github.com/brpratt/advent-of-code/file"
)

type point struct {
	x int
	y int
}

const (
	up = iota
	right
	down
	left
)

type robot struct {
	location point
	heading  int
}

func (r *robot) turnLeft() {
	if r.heading == up {
		r.heading = left
	} else {
		r.heading--
	}
}

func (r *robot) turnRight() {
	if r.heading == left {
		r.heading = up
	} else {
		r.heading++
	}
}

func (r *robot) move() {
	switch r.heading {
	case up:
		r.location.y++
	case right:
		r.location.x++
	case down:
		r.location.y--
	case left:
		r.location.x--
	}
}

func paint(program []int, start int) map[point]int {
	in := make(chan int, 1)
	out := make(chan int)
	c := intcode.NewComputer(program, in, out)

	go func() {
		err := c.Run()
		if err != nil {
			panic(err)
		}
	}()

	painted := make(map[point]int)
	r := robot{}

	painted[point{0, 0}] = start

	for {
		color, found := painted[r.location]
		if found {
			in <- color
		} else {
			in <- 0
		}

		color, more := <-out
		if !more {
			break
		}

		painted[r.location] = color

		rotation, more := <-out
		if !more {
			break
		}

		if rotation == 0 {
			r.turnLeft()
		} else {
			r.turnRight()
		}

		r.move()
	}

	return painted
}

func part01(program []int) int {
	painted := paint(program, 0)

	return len(painted)
}

func part02(program []int) string {
	var builder strings.Builder
	painted := paint(program, 1)

	maxX := 0
	minX := 0
	maxY := 0
	minY := 0

	for point := range painted {
		if point.x > maxX {
			maxX = point.x
		}
		if point.x < minX {
			minX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
		if point.y < minY {
			minY = point.y
		}
	}

	width := (maxX - minX) + 1
	height := (maxY - minY) + 1

	hull := make([][]int, height)
	for i := range hull {
		hull[i] = make([]int, width)
	}

	offX := -minX
	offY := -minY

	for point, color := range painted {
		hull[height-(point.y+offY+1)][point.x+offX] = color
	}

	for _, row := range hull {
		for _, color := range row {
			if color == 0 {
				fmt.Fprint(&builder, " ")
			} else {
				fmt.Fprint(&builder, "#")
			}
		}
		fmt.Fprintln(&builder)
	}

	return builder.String()
}

func main() {
	lines := file.Must(file.ReadLines("input.txt"))
	program := intcode.FromText(lines[0])

	fmt.Println(part01(program))
	fmt.Println(part02(program))
}
