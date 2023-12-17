package grid

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
