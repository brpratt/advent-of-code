package day15

import (
	"aoc2019/intcode"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	North = 1
	South = 2
	West  = 3
	East  = 4
)

const (
	Wall           = 0
	Moved          = 1
	MovedAndOxygen = 2
)

type Vector struct {
	x   int
	y   int
	dir int
}

type Droid struct {
	in  chan int
	out chan int
	c   intcode.Computer
	x   int
	y   int
	dir int
}

func NewDroid(program []int) *Droid {
	in := make(chan int)
	out := make(chan int)
	c := intcode.NewComputer(program, in, out)

	go c.Run()

	return &Droid{
		in,
		out,
		c,
		0,
		0,
		North,
	}
}

func (d *Droid) Move() int {
	d.in <- d.dir
	status := <-d.out
	
	if status == Moved || status == MovedAndOxygen {
		case d.dir == North:
			d.dir = East
		case d.dir == East:
			d.dir = South
		case d.dir == South:
			d.dir = West
		case d.dir == West:
			d.dir = North
	}
}

func (d *Droid) TurnRight() {
	switch {
	case d.dir == North:
		d.dir = East
	case d.dir == East:
		d.dir = South
	case d.dir == South:
		d.dir = West
	case d.dir == West:
		d.dir = North
	}
}

func SolvePart01(program []int) int {
	d := NewDroid(program)

	for {
		status := d.Move()
		fmt.Println(status)
		if status == Wall {
			d.TurnRight()
		}
		if status == MovedAndOxygen {
			break
		}
	}

	return 0
}

func SolvePart02(program []int) int {
	return 0
}

func Solve(part int, input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	values := strings.Split(scanner.Text(), ",")

	program := make([]int, len(values))
	for i, v := range values {
		num, _ := strconv.Atoi(v)
		program[i] = num
	}

	if part == 1 {
		return SolvePart01(program)
	}

	return SolvePart02(program)
}
