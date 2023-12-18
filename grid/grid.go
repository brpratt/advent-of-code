package grid

import "fmt"

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func Move(p Point, d Direction) Point {
	var np = p

	switch d {
	case Up:
		np.Y--
	case Right:
		np.X++
	case Down:
		np.Y++
	case Left:
		np.X--
	}

	return np
}

type Grid[T comparable] [][]T

func (g Grid[T]) InBound(p Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(g) && p.X < len(g[p.Y])
}

func (g Grid[T]) SetValue(p Point, t T) bool {
	if ok := g.InBound(p); ok {
		g[p.Y][p.X] = t
		return true
	}
	return false
}

func (g Grid[T]) Value(p Point) (T, bool) {
	var t T
	var ok bool

	if ok = g.InBound(p); ok {
		t = g[p.Y][p.X]
	}

	return t, ok
}

func (g Grid[T]) Find(t T) (Point, bool) {
	var p Point

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if g[y][x] == t {
				p.Y = y
				p.X = x
				return p, true
			}
		}
	}

	return p, false
}

func (g Grid[T]) FindAll(t T) []Point {
	points := make([]Point, 0)

	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			if g[y][x] == t {
				points = append(points, Point{Y: y, X: x})
			}
		}
	}

	return points
}

func (g Grid[T]) Row(n int) []T {
	if n < 0 || n >= len(g) {
		return nil
	}

	return g[n]
}

func (g Grid[T]) Column(n int) []T {
	if n < 0 || n >= len(g[0]) {
		return nil
	}

	c := make([]T, len(g))
	for i := 0; i < len(g); i++ {
		c[i] = g[i][n]
	}

	return c
}
