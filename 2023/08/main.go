package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/brpratt/advent-of-code/file"
)

var elementRegex *regexp.Regexp = regexp.MustCompile(`([A-Z1-9]+) = \(([A-Z1-9]+), ([A-Z1-9]+)\)`)

type network struct {
	instructions string
	pairs        map[string]struct {
		left  string
		right string
	}
}

func parseNetwork(r io.Reader) (n network) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		return
	}
	n.instructions = scanner.Text()
	scanner.Scan()

	n.pairs = make(map[string]struct {
		left  string
		right string
	})

	for scanner.Scan() {
		matches := elementRegex.FindStringSubmatch(scanner.Text())
		node := matches[1]
		n.pairs[node] = struct {
			left  string
			right string
		}{matches[2], matches[3]}
	}

	return
}

func networkCycle(network network, start string) (cycle int) {
	insts := []rune(network.instructions)
	instsindex := 0
	node := start

	for !strings.HasSuffix(node, "Z") {
		inst := insts[instsindex]
		instsindex++
		if instsindex >= len(insts) {
			instsindex = 0
		}

		if inst == 'L' {
			node = network.pairs[node].left
		} else {
			node = network.pairs[node].right
		}

		cycle++
	}

	return
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	network := parseNetwork(input)

	fmt.Println(part01(network))
	fmt.Println(part02(network))
}

func part01(network network) int {
	return networkCycle(network, "AAA")
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, xs ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(xs); i++ {
		result = lcm(result, xs[i])
	}

	return result
}

func part02(network network) int {
	cycles := make([]int, 0)

	for node := range network.pairs {
		if strings.HasSuffix(node, "A") {
			cycles = append(cycles, networkCycle(network, node))
		}
	}

	steps := lcm(cycles[0], cycles[1], cycles[2:]...)

	return steps
}
