package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/parse"
)

type equation struct {
	value    int
	operands []int
}

func parseEquation(line string) equation {
	var eq equation

	sides := strings.Split(line, ": ")
	eq.value = parse.MustAtoi(sides[0])
	rest := sides[1]

	for _, op := range strings.Split(rest, " ") {
		operand := parse.MustAtoi(op)
		eq.operands = append(eq.operands, operand)
	}

	return eq
}

func parseInput(r io.Reader) []equation {
	var equations []equation

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		equations = append(equations, parseEquation(line))
	}

	return equations
}

func ndigits(n int) int {
	var digits int

	for n > 0 {
		n /= 10
		digits++
	}

	return digits
}

func shift(n int, ndigits int) int {
	for ndigits > 0 {
		n *= 10
		ndigits--
	}

	return n
}

func equal(operands []int, operators []string, target int) bool {
	acc := operands[0]

	for i := 0; i < len(operands)-1; i++ {
		switch operators[i] {
		case "+":
			acc += operands[i+1]
		case "*":
			acc *= operands[i+1]
		case "||":
			acc = shift(acc, ndigits(operands[i+1])) + operands[i+1]
		}

		if acc > target {
			return false
		}
	}

	return acc == target
}

func spin(operators []string, index int, concat bool) bool {
	if index > len(operators)-1 {
		return false
	}

	current := operators[index]

	switch {
	case current == "+":
		operators[index] = "*"
	case current == "*" && concat:
		operators[index] = "||"
	case current == "*" && !concat:
		operators[index] = "+"
		return spin(operators, index+1, concat)
	case current == "||":
		operators[index] = "+"
		return spin(operators, index+1, concat)
	}

	return true
}

func hasSolution(eq equation, concat bool) bool {
	operators := make([]string, len(eq.operands)-1)
	for i := 0; i < len(operators); i++ {
		operators[i] = "+"
	}

	for {
		if equal(eq.operands, operators, eq.value) {
			return true
		}

		if !spin(operators, 0, concat) {
			return false
		}
	}
}

func part01(equations []equation) int {
	var sum int

	for _, eq := range equations {
		if hasSolution(eq, false) {
			sum += eq.value
		}
	}

	return sum
}

func part02(equations []equation) int {
	var sum int

	for _, eq := range equations {
		if hasSolution(eq, true) {
			sum += eq.value
		}
	}

	return sum
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	equations := parseInput(input)

	fmt.Println(part01(equations))
	fmt.Println(part02(equations))
}
