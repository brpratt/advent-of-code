package main

import (
	"fmt"

	"github.com/brpratt/advent-of-code/2019/intcode"
	"github.com/brpratt/advent-of-code/file"
)

func part01(program []int) int {
	in := make(chan int)
	out := make(chan int)
	c := intcode.NewComputer(program, in, out)

	go c.Run()

	in <- 1

	var result int

	for {
		r, more := <-out
		if !more {
			break
		}
		result = r
	}

	return result
}

func part02(program []int) int {
	in := make(chan int)
	out := make(chan int)
	c := intcode.NewComputer(program, in, out)

	go c.Run()
	in <- 2
	return <-out
}

func main() {
	lines := file.Must(file.ReadLines("input.txt"))
	program := intcode.FromText(lines[0])

	fmt.Println(part01(program))
	fmt.Println(part02(program))
}
