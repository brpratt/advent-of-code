package y2023

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
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

func SolveD08P01(r io.Reader) (string, error) {
	network := parseNetwork(r)

	steps := networkCycle(network, "AAA")

	return strconv.Itoa(steps), nil
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

func SolveD08P02(r io.Reader) (string, error) {
	network := parseNetwork(r)

	cycles := make([]int, 0)

	for node := range network.pairs {
		if strings.HasSuffix(node, "A") {
			cycles = append(cycles, networkCycle(network, node))
		}
	}

	steps := lcm(cycles[0], cycles[1], cycles[2:]...)

	return strconv.Itoa(steps), nil
}
