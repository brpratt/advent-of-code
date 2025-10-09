package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/brpratt/advent-of-code/file"
)

type point struct {
	x int
	y int
}

type direction int

const (
	up direction = iota
	right
	down
	left
)

type segment struct {
	direction direction
	count     int
}

type path []segment

func parseSegment(input string) segment {
	if len(input) < 2 {
		panic(fmt.Sprintf("malformed segment: %s", input))
	}

	var direction direction
	switch input[0] {
	case 'U':
		direction = up
	case 'R':
		direction = right
	case 'D':
		direction = down
	case 'L':
		direction = left
	default:
		panic(fmt.Sprintf("malformed segment: %s", input))
	}

	count, err := strconv.Atoi(input[1:])
	if err != nil {
		panic(fmt.Sprintf("malformed segment: %s", input))
	}

	return segment{direction, count}
}

func parsePath(input string) path {
	rawSegments := strings.Split(input, ",")
	path := make([]segment, len(rawSegments))

	for i, s := range rawSegments {
		path[i] = parseSegment(s)
	}

	return path
}

func segmentPoints(o point, s segment) (point, []point) {
	var pt point
	pts := make([]point, s.count)

	switch s.direction {
	case up:
		for i := 0; i < s.count; i++ {
			pts[i] = point{o.x, o.y + i + 1}
		}

		pt = point{o.x, o.y + s.count}
	case right:
		for i := 0; i < s.count; i++ {
			pts[i] = point{o.x + i + 1, o.y}
		}

		pt = point{o.x + s.count, o.y}
	case down:
		for i := 0; i < s.count; i++ {
			pts[i] = point{o.x, o.y - (i + 1)}
		}

		pt = point{o.x, o.y - s.count}
	case left:
		for i := 0; i < s.count; i++ {
			pts[i] = point{o.x - (i + 1), o.y}
		}

		pt = point{o.x - s.count, o.y}
	}

	return pt, pts
}

func points(o point, p path) []point {
	pointCount := 0
	for _, seg := range p {
		pointCount += seg.count
	}

	pts := make([]point, pointCount)
	i := 0

	for _, segment := range p {
		pt, segPts := segmentPoints(o, segment)

		o = pt

		for j := range segPts {
			pts[i] = segPts[j]
			i++
		}
	}

	return pts
}

func toSet(points []point) map[point]bool {
	ps := make(map[point]bool)

	for _, p := range points {
		if !ps[p] {
			ps[p] = true
		}
	}

	return ps
}

func steps(points []point) map[point]int {
	steps := make(map[point]int)

	for i, p := range points {
		if steps[p] == 0 {
			steps[p] = i + 1
		}
	}

	return steps
}

func intersect(steps1 map[point]int, steps2 map[point]int) map[point]bool {
	common := make(map[point]bool)

	for point := range steps1 {
		if steps2[point] != 0 {
			common[point] = true
		}
	}

	return common
}

func distance(o point, p point) int {
	return int(math.Abs(float64(o.x-p.x))) +
		int(math.Abs(float64(o.y-p.y)))
}

func min(vars ...int) int {
	min := vars[0]

	for _, v := range vars {
		if v < min {
			min = v
		}
	}

	return min
}

func part01(input []string) int {
	path1 := parsePath(input[0])
	path2 := parsePath(input[1])
	o := point{0, 0}

	points1 := points(o, path1)
	points2 := points(o, path2)

	steps1 := steps(points1)
	steps2 := steps(points2)

	commonPoints := intersect(steps1, steps2)

	distances := make([]int, len(commonPoints))
	i := 0

	for point := range commonPoints {
		distances[i] = distance(o, point)
		i++
	}

	return min(distances...)
}

func part02(input []string) int {
	path1 := parsePath(input[0])
	path2 := parsePath(input[1])
	o := point{0, 0}

	points1 := points(o, path1)
	points2 := points(o, path2)

	steps1 := steps(points1)
	steps2 := steps(points2)

	commonPoints := intersect(steps1, steps2)
	combinedSteps := make([]int, len(commonPoints))

	i := 0
	for point := range commonPoints {
		combinedSteps[i] = steps1[point] + steps2[point]
		i++
	}

	return min(combinedSteps...)
}

func main() {
	lines := file.Must(file.ReadLines("input.txt"))

	fmt.Println(part01(lines))
	fmt.Println(part02(lines))
}
