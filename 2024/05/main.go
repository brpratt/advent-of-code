package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/parse"
)

type pageOrder struct {
	before int
	after  int
}

type update []int

func parseInput(r io.Reader) ([]pageOrder, []update) {
	pageOrders := make([]pageOrder, 0)
	updates := make([]update, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		var po pageOrder
		fmt.Sscanf(line, "%d|%d", &po.before, &po.after)
		pageOrders = append(pageOrders, po)
	}

	for scanner.Scan() {
		line := scanner.Text()
		var update update

		for _, num := range strings.Split(line, ",") {
			update = append(update, parse.MustAtoi(num))
		}

		updates = append(updates, update)
	}

	return pageOrders, updates
}

func isCorrectOrder(update update, pageOrders []pageOrder) bool {
	for i := range update {
		for j := 0; j < i; j++ {
			for _, po := range pageOrders {
				if po.before == update[i] && po.after == update[j] {
					return false
				}
			}
		}
		for j := i + 1; j < len(update); j++ {
			for _, po := range pageOrders {
				if po.before == update[j] && po.after == update[i] {
					return false
				}
			}
		}
	}

	return true
}

func part01(pageOrders []pageOrder, updates []update) int {
	var sum int

	for _, update := range updates {
		if isCorrectOrder(update, pageOrders) {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func part02(pageOrders []pageOrder, updates []update) int {
	unordered := make([]update, 0)

	for _, update := range updates {
		if !isCorrectOrder(update, pageOrders) {
			unordered = append(unordered, update)
		}
	}

	var sum int

	for _, update := range unordered {
		sort.Slice(update, func(a, b int) bool {
			for _, po := range pageOrders {
				if po.before == update[a] && po.after == update[b] {
					return true
				}
			}

			return false
		})

		sum += update[len(update)/2]
	}

	return sum
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	pageOrders, updates := parseInput(input)

	fmt.Println(part01(pageOrders, updates))
	fmt.Println(part02(pageOrders, updates))
}
