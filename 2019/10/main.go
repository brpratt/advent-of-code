package day10

import (
	"bufio"
	"io"
	"math"
	"sort"
)

type point struct {
	x int
	y int
}

type pointangle struct {
	x     int
	y     int
	angle float64
}

type slope struct {
	rise int
	run  int
}

func extractPoints(input []string) []point {
	points := make([]point, 0)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == '#' {
				points = append(points, point{j, i})
			}
		}
	}
	return points
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func sign(x int) int {
	if x < 0 {
		return -1
	}

	return 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// manhattan
func distanceBetween(p0, p1 point) int {
	return abs(p1.x-p0.x) + abs(p1.y-p0.y)
}

func slopeBetween(a, b point) slope {
	rise := b.y - a.y
	run := b.x - a.x

	if rise == 0 && run == 0 {
		return slope{0, 0}
	}

	if rise == 0 {
		return slope{0, sign(run)}
	}

	if run == 0 {
		return slope{sign(rise), 0}
	}

	n := gcd(abs(rise), abs(run))
	if n == 1 {
		return slope{rise, run}
	}

	return slope{rise / n, run / n}
}

func laserAngle(s slope) float64 {
	a := math.Atan(float64(s.rise)/float64(s.run)) * 180 / math.Pi

	if s.run >= 0 {
		return a + 90
	}

	return a + 270
}

func visible(p0 point, points []point) int {
	m := make(map[slope]bool)
	for _, p1 := range points {
		if p0 == p1 {
			continue
		}

		s := slopeBetween(p0, p1)
		m[s] = true
	}

	return len(m)
}

func bestLocation(points []point) (point, int) {
	var maxPoint point
	maxVisible := 0

	for _, p0 := range points {
		v := visible(p0, points)
		if v > maxVisible {
			maxPoint = p0
			maxVisible = v
		}
	}

	return maxPoint, maxVisible
}

// first slice contains the visible points
// second slice contains the remaining points
func extractVisible(p0 point, points []point) ([]point, []point) {
	nonvpoints := make([]point, 0)
	slopes := make(map[slope]point)
	for _, p1 := range points {
		if p0 == p1 {
			continue
		}

		s := slopeBetween(p0, p1)
		p2, ok := slopes[s]

		if !ok {
			slopes[s] = p1
		}

		if distanceBetween(p0, p1) < distanceBetween(p0, p2) {
			nonvpoints = append(nonvpoints, p2)
			slopes[s] = p1
		}
	}

	vpoints := make([]point, 0, len(slopes))
	for _, p := range slopes {
		vpoints = append(vpoints, p)
	}

	return vpoints, nonvpoints
}

func addLaserAngle(p0 point, points []point) []pointangle {
	pointangles := make([]pointangle, 0, len(points))

	for _, p1 := range points {
		angle := laserAngle(slopeBetween(p0, p1))
		pointangles = append(pointangles, pointangle{p1.x, p1.y, angle})
	}

	return pointangles
}

func SolvePart01(input []string) int {
	points := extractPoints(input)

	_, count := bestLocation(points)

	return count
}

func SolvePart02(input []string, count int) int {
	points := extractPoints(input)
	p0, _ := bestLocation(points)
	destroyed := 0

	for {
		vpoints, nonvpoints := extractVisible(p0, points)

		pointangles := addLaserAngle(p0, vpoints)

		sort.Slice(pointangles, func(i, j int) bool {
			return pointangles[i].angle < pointangles[j].angle
		})

		for _, pointangle := range pointangles {
			destroyed++

			if destroyed == count {
				return pointangle.x*100 + pointangle.y
			}
		}

		points = nonvpoints
	}
}

func Solve(part int, input io.Reader) int {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if part == 1 {
		return SolvePart01(lines)
	}

	return SolvePart02(lines, 200)
}
