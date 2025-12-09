package main

import (
	"fmt"

	"github.com/brpratt/advent-of-code/grid"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func parsePoint(input string) grid.Point {
	var x, y int
	fmt.Sscanf(input, "%d,%d", &x, &y)
	return grid.Point{X: x, Y: y}
}

func parsePoints(lines []string) []grid.Point {
	points := make([]grid.Point, len(lines))
	for i, line := range lines {
		points[i] = parsePoint(line)
	}
	return points
}

func area(a, b grid.Point) int {
	height := abs(a.Y-b.Y) + 1
	width := abs(a.X-b.X) + 1

	return height * width
}

func part01(points []grid.Point) int {
	var max int

	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			area := area(points[i], points[j])
			if area > max {
				max = area
			}
		}
	}

	return max
}
