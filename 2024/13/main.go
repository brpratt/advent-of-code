package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/grid"
)

type game struct {
	buttonA grid.Point
	buttonB grid.Point
	prize   grid.Point
}

func parseInput(r io.Reader) []game {
	games := make([]game, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var g game

		fmt.Sscanf(scanner.Text(), "Button A: X+%d, Y+%d", &g.buttonA.X, &g.buttonA.Y)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Button B: X+%d, Y+%d", &g.buttonB.X, &g.buttonB.Y)
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "Prize: X=%d, Y=%d", &g.prize.X, &g.prize.Y)
		scanner.Scan()

		games = append(games, g)
	}

	return games
}

func solve(a1, b1, c1, a2, b2, c2 int) (int, int) {
	determinant := a1*b2 - a2*b1
	if determinant == 0 {
		return 0, 0
	}

	x := (c1*b2 - c2*b1) / determinant
	y := (a1*c2 - a2*c1) / determinant

	return x, y
}

type presses struct {
	a int
	b int
}

func play(g game) presses {
	a, b := solve(g.buttonA.X, g.buttonB.X, g.prize.X, g.buttonA.Y, g.buttonB.Y, g.prize.Y)

	if (g.buttonA.X*a)+(g.buttonB.X*b) != g.prize.X {
		return presses{a: 0, b: 0}
	}
	if (g.buttonA.Y*a)+(g.buttonB.Y*b) != g.prize.Y {
		return presses{a: 0, b: 0}
	}

	return presses{a: a, b: b}
}

func cost(c presses) int {
	return (c.a * 3) + c.b
}

func part01(games []game) int {
	var sum int

	for _, game := range games {
		combination := play(game)
		sum += cost(combination)
	}

	return sum
}

func part02(games []game) int {
	var sum int

	for _, game := range games {
		game.prize.X += 10000000000000
		game.prize.Y += 10000000000000
		combination := play(game)
		sum += cost(combination)
	}

	return sum
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	games := parseInput(input)

	fmt.Println(part01(games))
	fmt.Println(part02(games))
}
