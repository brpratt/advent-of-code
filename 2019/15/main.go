package main

import (
	"fmt"

	"github.com/brpratt/advent-of-code/2019/intcode"
	"github.com/brpratt/advent-of-code/file"
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

	return 0
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

func part01(program []int) int {
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

func part02(program []int) int {
	return 0
}

func main() {
	lines := file.Must(file.ReadLines("input.txt"))
	program := intcode.FromText(lines[0])

	fmt.Println(part01(program))
	fmt.Println(part02(program))
}
