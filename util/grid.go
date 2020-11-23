package util

import (
	"bytes"
	"sync"
)

// Grid represents an array of points.
type Grid struct {
	sync.RWMutex
	points []point
}

// point represnts a cartesian point.
type point struct {
	X, Y int
	V    string
}

// Size ...
func (g *Grid) Size() (int, int) {
	var minX, minY, maxX, maxY int
	g.RLock()
	defer g.RUnlock()
	for _, p := range g.points {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return (maxX - minX) + 1, (maxY - minY) + 1
}

// String ...
func (g *Grid) String() string {
	x, y := g.Size()
	graph := [][]byte{}
	for i := 0; i < x; i++ {
		row := []byte{}
		for j := 0; j < y; j++ {
			row = append(row, '.')
		}
		graph = append(graph, row)
	}
	g.RLock()
	defer g.RUnlock()
	trans := func(p, q int) (int, int) {
		return p + x, q + y
	}
	for _, p := range g.points {
		x, y := trans(p.X, p.Y)
		graph[x][y] = '#'
	}
	return string(bytes.Join(graph, []byte("\n")))
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
func (g *Grid) AddPoint(x, y int, v string) {
	g.Lock()
	defer g.Unlock()
	g.points = append(g.points, point{X: x, Y: y, V: v})
}

// AddStrip ...
func (g *Grid) AddStrip(x, y, steps int, dir rune, v string) {
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
		g.AddPoint(x, y, v)
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
