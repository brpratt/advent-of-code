package y2019

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/brpratt/advent-of-code/y2019/intcode"
)

func readValues(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}

func toIntcode(input []string) []int {
	intcode := make([]int, len(input))

	for i, v := range input {
		x, _ := strconv.Atoi(v)
		intcode[i] = x
	}

	return intcode
}

func SolveD02P01(r io.Reader) (string, error) {
	values := readValues(r)
	program := toIntcode(values)

	c := intcode.NewComputer(program, make(chan int), make(chan int))
	c.Memory[1] = 12
	c.Memory[2] = 2
	c.Run()

	return strconv.Itoa(c.Memory[0]), nil
}

func SolveD02P02(r io.Reader) (string, error) {
	values := readValues(r)
	program := toIntcode(values)

	for noun := 0; noun <= 100; noun++ {
		for verb := 0; verb <= 100; verb++ {
			c := intcode.NewComputer(program, make(chan int), make(chan int))
			c.Memory[1] = noun
			c.Memory[2] = verb
			c.Run()

			if c.Memory[0] == 19690720 {
				result := 100*noun + verb
				return strconv.Itoa(result), nil
			}
		}
	}

	return "", errors.New("could not find solution")
}
