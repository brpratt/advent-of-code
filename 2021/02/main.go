package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brpratt/advent-of-code/file"
)

type command struct {
	direction string
	amount    int
}

func main() {
	lines := file.Must(file.ReadLines("input.txt"))

	commands := make([]command, 0)
	for _, line := range lines {
		segments := strings.Split(line, " ")
		amount, err := strconv.Atoi(segments[1])
		if err != nil {
			panic(err)
		}

		commands = append(commands, command{segments[0], amount})
	}

	solvePart01(commands)
	solvePart02(commands)
}

func solvePart01(commands []command) {
	var pos struct {
		horizontal int
		depth      int
	}

	for _, command := range commands {
		switch command.direction {
		case "forward":
			pos.horizontal += command.amount
		case "down":
			pos.depth += command.amount
		case "up":
			pos.depth -= command.amount
		}
	}

	fmt.Println(pos.depth * pos.horizontal)
}

func solvePart02(commands []command) {
	var pos struct {
		aim        int
		horizontal int
		depth      int
	}

	for _, command := range commands {
		switch command.direction {
		case "forward":
			pos.horizontal += command.amount
			pos.depth += pos.aim * command.amount
		case "down":
			pos.aim += command.amount
		case "up":
			pos.aim -= command.amount
		}
	}

	fmt.Println(pos.depth * pos.horizontal)
}
