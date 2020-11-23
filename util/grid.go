package util

import "sync"

// Grid represents a NxN grid of points.
type Grid struct {
	sync.RWMutex
	Point [][]string
}

// Resize returns a new grid size with size of the positive X axis containing
// all of the source grid's points.
func (g *Grid) Resize(size int) {
	if len(g.Point) >= size {
		return
	}
	g.Lock()
	defer g.Unlock()
	r := make([][]string, size)
	for j := 0; j < size; j++ {
		newRow := make([]string, size)
		r[j] = newRow
	}
	for j, row := range g.Point {
		for i, x := range row {
			r[j][i] = x
		}
	}
	g.Point = r
}

// String ...
func (g *Grid) String() string {
	g.RLock()
	defer g.RUnlock()
	r := ""
	for i := len(g.Point) - 1; i >= 0; i-- {
		for _, x := range g.Point[i] {
			if x == "" {
				x = "."
			}
			r += x
		}
		r += "\n"
	}
	return r
}

// Translate returns Cartesian (-N,N) coordinates translated into grid coordinates [0, N).
func (*Grid) Translate(x, y int) (int, int) {
	if x > 0 {
		x = 2 * x
	} else {
		x = -x
	}

	if y > 0 {
		y = 2 * y
	} else {
		y = -y
	}

	return x, y
}

// Return the Cartesian point walking from (X, Y) steps in direction.
func (*Grid) Walk(x, y, steps int, dir rune) (int, int) {
	switch dir {
	case 'U':
		return x, y + steps
	case 'D':
		return x, y - steps
	case 'L':
		return x - steps, y
	case 'R':
		return x + steps, y
	}
	return x, y
}

// AddPoint ...
func (g *Grid) AddPoint(x, y int, z string) {
	max := x
	if y > max {
		max = y
	}
	g.Resize(1 + max)
	g.Lock()
	defer g.Unlock()
	g.Point[y][x] = z
}

// AddStrip ...
func (g *Grid) AddStrip(x, y, steps int, dir rune, z string) {
	var addX, addY int
	switch dir {
	case 'U':
		addY = 1
	case 'D':
		addY = -1
	case 'L':
		addX = -1
	case 'R':
		addX = 1
	}
	for i := 0; i < steps; i++ {
		g.AddPoint(x, y, z)
		x += addX
		y += addY
	}
}

// Distance returns the taxicab distance between two points.
func (*Grid) Distance(x1, y1, x2, y2 int) int {
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
