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
	for j := 0; j < size; j++ {
		newRow := make([]string, size)
		r.Point[j] = newRow
	}
	for j, row := range g.Point {
		for i, x := range row {
			r.Point[j][i] = x
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

// Translate returns Cartesian (-N,N) coordinates translated into grid coordinates [0, N).
func (g Grid) Translate(x, y int) (int, int) {
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
func (g Grid) Walk(x, y, steps int, dir rune) (int, int) {
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
}

// AddPoint ...
func (g Grid) AddPoint(x, y int, z string) *Grid {
	max := x
	if y > max {
		max = y
	}
	r := g.Resize(1 + max)
	r.Point[y][x] = z
	return r
}

// AddStrip ...
func (g Grid) AddStrip(x, y, steps int, dir rune, z string) *Grid {
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
	r := &g
	for i := 0; i < steps; i++ {
		r = r.AddPoint(x, y, z)
		x += addX
		y += addY
	}
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
