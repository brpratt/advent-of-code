package main

import (
	"fmt"
	"math"
	"slices"
)

type box struct {
	x int
	y int
	z int
}

type edge struct {
	from     box
	to       box
	distance float64
}

func newEdge(from, to box) edge {
	distance := math.Sqrt(
		math.Pow(float64(from.x-to.x), 2) +
			math.Pow(float64(from.y-to.y), 2) +
			math.Pow(float64(from.z-to.z), 2))

	return edge{
		from:     from,
		to:       to,
		distance: distance,
	}
}

type minheap struct {
	edges []edge
}

func downidx(mh *minheap, idx int) int {
	left := (idx * 2) + 1
	right := (idx * 2) + 2

	min := idx

	if left < len(mh.edges) && mh.edges[left].distance < mh.edges[min].distance {
		min = left
	}

	if right < len(mh.edges) && mh.edges[right].distance < mh.edges[min].distance {
		min = right
	}

	if min != idx {
		return min
	}

	return -1
}

func (mh *minheap) pop() edge {
	if len(mh.edges) == 0 {
		return edge{}
	}

	top := mh.edges[0]
	mh.edges[0] = mh.edges[len(mh.edges)-1]
	mh.edges = mh.edges[:len(mh.edges)-1]

	if len(mh.edges) == 0 {
		return top
	}

	curridx := 0
	swapidx := downidx(mh, curridx)

	for swapidx != -1 {
		mh.edges[curridx], mh.edges[swapidx] = mh.edges[swapidx], mh.edges[curridx]
		curridx = swapidx
		swapidx = downidx(mh, curridx)
	}

	return top
}

func (mh *minheap) add(e edge) {
	mh.edges = append(mh.edges, e)

	curridx := len(mh.edges) - 1
	previdx := (curridx - 1) / 2

	for curridx != 0 && mh.edges[curridx].distance < mh.edges[previdx].distance {
		mh.edges[curridx], mh.edges[previdx] = mh.edges[previdx], mh.edges[curridx]
		curridx = previdx
		previdx = (curridx - 1) / 2
	}
}

type circuit struct {
	boxes map[box]bool
	edges []edge
}

func newCircuit() circuit {
	return circuit{
		boxes: make(map[box]bool),
		edges: make([]edge, 0),
	}
}

func (c *circuit) add(e edge) {
	c.boxes[e.from] = true
	c.boxes[e.to] = true
	c.edges = append(c.edges, e)
}

func (c *circuit) size() int {
	return len(c.boxes)
}

type circuitset struct {
	circuits []circuit
}

func newCircuitSet() circuitset {
	return circuitset{circuits: make([]circuit, 0)}
}

func (cs *circuitset) add(e edge) {
	matching := make([]int, 0)

	for i, circuit := range cs.circuits {
		if circuit.boxes[e.from] || circuit.boxes[e.to] {
			matching = append(matching, i)
		}
	}

	if len(matching) == 0 {
		circuit := newCircuit()
		circuit.add(e)
		cs.circuits = append(cs.circuits, circuit)
		return
	}

	if len(matching) == 1 {
		circuit := &cs.circuits[matching[0]]
		circuit.add(e)
		return
	}

	circuit := &cs.circuits[matching[0]]
	oldcircuit := &cs.circuits[matching[1]]

	for _, oldedge := range oldcircuit.edges {
		circuit.add(oldedge)
	}

	circuit.add(e)
	cs.circuits = slices.Delete(cs.circuits, matching[1], matching[1]+1)
}

func parseBox(line string) box {
	var x, y, z int
	fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
	return box{x: x, y: y, z: z}
}

func parseBoxes(lines []string) []box {
	boxes := make([]box, len(lines))
	for i := range lines {
		boxes[i] = parseBox(lines[i])
	}
	return boxes
}

func part01(boxes []box, connections int) int {
	heap := minheap{edges: make([]edge, 0)}

	for i := 0; i < len(boxes)-1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			heap.add(newEdge(boxes[i], boxes[j]))
		}
	}

	cs := newCircuitSet()

	for range connections {
		edge := heap.pop()
		cs.add(edge)
	}

	largest := make([]int, 3)

	for i := range largest {
		largestidx := 0
		for checkidx := 1; checkidx < len(cs.circuits); checkidx++ {
			if cs.circuits[checkidx].size() > cs.circuits[largestidx].size() {
				largestidx = checkidx
			}
		}

		largest[i] = cs.circuits[largestidx].size()
		cs.circuits = slices.Delete(cs.circuits, largestidx, largestidx+1)
	}

	return largest[0] * largest[1] * largest[2]
}

func part02(boxes []box) int {
	heap := minheap{edges: make([]edge, 0)}

	for i := 0; i < len(boxes)-1; i++ {
		for j := i + 1; j < len(boxes); j++ {
			heap.add(newEdge(boxes[i], boxes[j]))
		}
	}

	cs := newCircuitSet()
	var edge edge

	for {
		if len(cs.circuits) == 1 && cs.circuits[0].size() == len(boxes) {
			break
		}

		edge = heap.pop()
		cs.add(edge)
	}

	return edge.from.x * edge.to.x
}
