package main

import (
	"fmt"
	"strings"

	"github.com/brpratt/advent-of-code/file"
)

func buildOrbitMap(input []string) map[string]string {
	m := make(map[string]string)

	for _, line := range input {
		segments := strings.Split(line, ")")
		m[segments[1]] = segments[0]
	}

	return m
}

func buildReverseOrbitMap(to string, m map[string]string) map[string]string {
	revM := make(map[string]string)
	next := m[to]

	for next != "COM" {
		revM[next] = to
		to = next
		next = m[next]
	}

	revM[next] = to

	return revM
}

func countOrbits(obj string, m map[string]string) int {
	count := 0

	for obj != "COM" {
		obj = m[obj]
		count++
	}

	return count
}

func countTransfers(from string, to string, m map[string]string) int {
	rev := buildReverseOrbitMap(to, m)
	curr := m[from]
	count := 0

	for _, ok := rev[curr]; !ok; _, ok = rev[curr] {
		curr = m[curr]
		count++
	}

	for rev[curr] != to {
		curr = rev[curr]
		count++
	}

	return count
}

func part01(orbits []string) int {
	m := buildOrbitMap(orbits)
	count := 0

	for key := range m {
		count += countOrbits(key, m)
	}

	return count
}

func part02(orbits []string) int {
	m := buildOrbitMap(orbits)

	return countTransfers("YOU", "SAN", m)
}

func main() {
	orbits := file.Must(file.ReadLines("input.txt"))

	fmt.Println(part01(orbits))
	fmt.Println(part02(orbits))
}
