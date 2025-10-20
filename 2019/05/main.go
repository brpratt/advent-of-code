package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brpratt/advent-of-code/2019/intcode"
	"github.com/brpratt/advent-of-code/file"
)

func toIntcode(line string) []int {
	values := strings.Split(line, ",")
	program := make([]int, len(values))
	for i, v := range values {
		num, _ := strconv.Atoi(v)
		program[i] = num
	}

	return program
}

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
	in <- 5
	return <-out
}

func main() {
	lines := file.Must(file.ReadLines("input.txt"))
	program := toIntcode(lines[0])

	fmt.Println(part01(program))
	fmt.Println(part02(program))
}
