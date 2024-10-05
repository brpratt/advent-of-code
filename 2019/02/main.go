package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brpratt/advent-of-code/2019/intcode"
	"github.com/brpratt/advent-of-code/file"
)

func main() {
	lines := file.Must(file.ReadLines("input.txt"))
	values := strings.Split(lines[0], ",")
	program := toProgram(values)

	fmt.Println(part01(program))
	fmt.Println(part02(program))
}

func toProgram(input []string) []int {
	intcode := make([]int, len(input))

	for i, v := range input {
		x, _ := strconv.Atoi(v)
		intcode[i] = x
	}

	return intcode
}

func part01(program []int) int {
	c := intcode.NewComputer(program, make(chan int), make(chan int))
	c.Memory[1] = 12
	c.Memory[2] = 2
	c.Run()

	return c.Memory[0]
}

func part02(program []int) int {
	for noun := 0; noun <= 100; noun++ {
		for verb := 0; verb <= 100; verb++ {
			c := intcode.NewComputer(program, make(chan int), make(chan int))
			c.Memory[1] = noun
			c.Memory[2] = verb
			c.Run()

			if c.Memory[0] == 19690720 {
				result := 100*noun + verb
				return result
			}
		}
	}

	panic("could not find solution")
}
