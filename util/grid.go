package util

// Grid represents a NxN grid of points.
type Grid struct {
	Point [][]string
}

// Resize returns a new grid size with size of the positive X axis containing
// all of the source grid's points.
func (g *Grid) Resize(size int) *Grid {
	if len(g.Point) >= size {
		return g
	}
	r := Grid{make([][]string, size)}
	for i := 0; i < size; i++ {
		newRow := make([]string, size)
		r.Point[i] = newRow
	}
	for i, row := range g.Point {
		for j, x := range row {
			r.Point[i][j] = x
		}
	}
	return &r
}

// String ...
func (g Grid) String() string {
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

// AddPoint ...
func (g Grid) AddPoint(x, y int, z string) *Grid {
	max := x
	if y > max {
		max = y
	}
	r := g.Resize(1 + max)
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
