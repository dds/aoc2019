package util

import (
	"github.com/pkg/math"
)

// Grid represents a NxN grid of points.
type Grid struct {
	Point [][]string
}

// Resize returns a new grid of size N containing all of the existing grids
// points.
func (g Grid) Resize(size int) Grid {
	if size <= len(g.Point) {
		return g
	}
	r := Grid{make([][]string, size)}
	for i, row := range g.Point {
		newRow := make([]string, size)
		copy(newRow, row)
		r.Point[i] = newRow
	}
	return r
}

// AddPoint ...
func (g Grid) AddPoint(x, y int, z string) Grid {
	var r Grid
	if x > len(g.Point) || y > len(g.Point) {
		r = g.Resize(math.MaxInt(x, y))
	}
	r.Point[x][y] = z
	return r
}

// Distance returns the taxicab distance between two points.
func Distance(x1, y1, x2, y2 int) int {
	var p1, p2 int
	if x2 > x1 {
		p1 = x2 - x1
	} else {
		p1 = x1 - x2
	}
	if y2 > y1 {
		p2 = y2 - y1
	} else {
		p2 = y1 - y2
	}
	return p1 + p2
}
