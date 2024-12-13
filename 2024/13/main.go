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

type combination struct {
	buttonA int
	buttonB int
}

func solve(g game) []combination {
	var (
		cache map[grid.Point]map[combination]struct{}
		iter  func(p grid.Point) map[combination]struct{}
	)

	cache = make(map[grid.Point]map[combination]struct{})

	iter = func(p grid.Point) map[combination]struct{} {
		if cached, ok := cache[p]; ok {
			return cached
		}

		if p == g.prize {
			set := make(map[combination]struct{})
			set[combination{buttonA: 0, buttonB: 0}] = struct{}{}
			return set
		}

		if p.X > g.prize.X || p.Y > g.prize.Y {
			return make(map[combination]struct{})
		}

		posA := grid.Point{
			X: p.X + g.buttonA.X,
			Y: p.Y + g.buttonA.Y,
		}

		posB := grid.Point{
			X: p.X + g.buttonB.X,
			Y: p.Y + g.buttonB.Y,
		}

		combosA := iter(posA)
		combosB := iter(posB)

		combinations := make(map[combination]struct{})
		for combo := range combosA {
			combo.buttonA++
			combinations[combo] = struct{}{}
		}
		for combo := range combosB {
			combo.buttonB++
			combinations[combo] = struct{}{}
		}

		cache[p] = combinations
		return combinations
	}

	set := iter(grid.Point{X: 0, Y: 0})
	combinations := make([]combination, 0)
	for combo := range set {
		combinations = append(combinations, combo)
	}

	return combinations
}

func cost(c combination) int {
	return (c.buttonA * 3) + c.buttonB
}

func part01(games []game) int {
	var sum int

	for _, game := range games {
		combinations := solve(game)
		if len(combinations) == 0 {
			continue
		}

		cheapest := cost(combinations[0])
		for i := 1; i < len(combinations); i++ {
			c := cost(combinations[i])
			if c < cheapest {
				cheapest = c
			}
		}

		sum += cheapest
	}

	return sum
}

func part02(games []game) int {
	var sum int

	for _, game := range games {
		game.prize.X += 10000000000000
		game.prize.Y += 10000000000000

		combinations := solve(game)
		if len(combinations) == 0 {
			continue
		}

		cheapest := cost(combinations[0])
		for i := 1; i < len(combinations); i++ {
			c := cost(combinations[i])
			if c < cheapest {
				cheapest = c
			}
		}

		sum += cheapest
	}

	return sum
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	stones := parseInput(input)

	fmt.Println(part01(stones))
	fmt.Println(part02(stones))
}
