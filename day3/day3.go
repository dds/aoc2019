package main

import (
	"bytes"
	"fmt"
	"strconv"

	pkgmath "github.com/pkg/math"
	"gonum.org/v1/gonum/spatial/kdtree"

	"github.com/dds/aoc2019/util"
)

var Input = util.ParseInput(util.Inputs[3], util.CSVParser)

func main() {
	path1, path2 := Input[0], Input[1]
	fmt.Println("part1", Cross(path1, path2))
	fmt.Println("part2", MinWalk(path1, path2))
}

func MinWalk(path1 []string, path2 []string) int {
	points := make(Points, 0)
	var x, y, step int = 0, 0, 1
	for _, i := range path1 {
		dir, paces := Parse(i)
		for _, p := range Walk(x, y, paces, dir) {
			p.steps = step
			step += 1
			points = append(points, p)
			x, y = p.x, p.y
		}
	}
	tree := kdtree.New(points, true)
	x, y, step = 0, 0, 1
	crosses := make([][2]Point, 0)
	for _, j := range path2 {
		dir, paces := Parse(j)
		for _, p := range Walk(x, y, paces, dir) {
			p.steps = step
			step += 1
			x, y = p.x, p.y
			q, d := tree.Nearest(p)
			if d == 0 {
				crosses = append(crosses, [2]Point{q.(Point), p})
			}
		}
	}
	if len(crosses) == 0 {
		return 0
	}
	minstep := crosses[0][0].steps + crosses[0][1].steps
	for _, t := range crosses[1:] {
		s := t[0].steps + t[1].steps
		if s < minstep {
			minstep = s
		}
	}
	return minstep
}

func Parse(pos string) (rune, int) {
	dir := pos[0]
	paces, err := strconv.Atoi(pos[1:])
	if err != nil {
		panic(err)
	}
	return rune(dir), paces
}

func Walk(x, y, paces int, dir rune) (r []Point) {
	var p, q int
	switch dir {
	case 'U':
		q = 1
	case 'D':
		q = -1
	case 'L':
		p = -1
	case 'R':
		p = 1
	}
	for i := 0; i < paces; i++ {
		pt := Point{x: x + p, y: y + q}
		x, y = pt.x, pt.y
		r = append(r, pt)
	}
	return
}

func Graph(tree *kdtree.Tree) string {
	fmt.Println(tree.Root, tree.Root)
	min, max := tree.Root.Min.(Point), tree.Root.Max.(Point)
	fmt.Println("min, max", min, max)
	width := max.x - min.x + 1
	height := max.y - min.y + 1
	fmt.Println("width, height", width, height)
	cells := make([][]byte, height)
	for i := 0; i < height; i++ {
		cells[i] = make([]byte, width)
		for j := 0; j < width; j++ {
			cells[i][j] = '.'
		}
	}
	f := func(c kdtree.Comparable, _ *kdtree.Bounding, _ int) bool {
		p := c.(Point)
		y := height - 1*(p.y-min.y) - 1
		x := p.x - min.x
		if y < 0 || x < 0 {
			panic(fmt.Errorf("p=%v, x=%v, y=%v", p, x, y))
		}
		cells[y][x] = '#'
		return false
	}
	tree.Do(f)
	return string(bytes.Join(cells, []byte("\n")))
}

func Cross(path1 []string, path2 []string) int {
	points := make(Points, 0)
	var x, y int
	for _, i := range path1 {
		dir, paces := Parse(i)
		points = append(points, Walk(x, y, paces, dir)...)
		x, y = int(points[len(points)-1].x), int(points[len(points)-1].y)
	}
	tree := kdtree.New(points, true)
	x, y = 0, 0
	crosses := make([]Point, 0)
	for _, j := range path2 {
		dir, paces := Parse(j)
		for _, p := range Walk(x, y, paces, dir) {
			x, y = p.x, p.y
			q, d := tree.Nearest(p)
			if d == 0 {
				crosses = append(crosses, q.(Point))
			}
		}
	}
	if len(crosses) == 0 {
		return 0
	}
	min := manhattan(0, 0, crosses[0].x, crosses[0].y)
	for _, t := range crosses {
		d := manhattan(0, 0, t.x, t.y)
		if d < min {
			min = d
		}
	}
	return min
}

func manhattan(x1, y1, x2, y2 int) int {
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

type Point struct {
	x, y  int
	steps int
}

func (p Point) Compare(c kdtree.Comparable, d kdtree.Dim) float64 {
	q := c.(Point)
	switch d {
	case 0:
		return float64(p.x - q.x)
	case 1:
		return float64(p.y - q.y)
	default:
		panic("extra dimension")
	}
}

func (p Point) Dims() int { return 2 }

func (p Point) Distance(c kdtree.Comparable) float64 {
	q := c.(Point)
	return float64(manhattan(q.x, q.y, p.x, p.y))
}

func (p Point) Extend(b *kdtree.Bounding) *kdtree.Bounding {
	if b == nil {
		b = &kdtree.Bounding{Point{x: p.x, y: p.y}, Point{x: p.x, y: p.y}}
	}
	min := b.Min.(Point)
	max := b.Max.(Point)
	min.x = pkgmath.MinInt(min.x, p.x)
	min.y = pkgmath.MinInt(min.y, p.y)
	max.x = pkgmath.MaxInt(max.x, p.x)
	max.y = pkgmath.MaxInt(max.y, p.y)
	*b = kdtree.Bounding{Min: min, Max: max}
	return b

}

type Points []Point

func (p Points) Bounds() *kdtree.Bounding {
	if len(p) == 0 {
		return nil
	}
	min, max := Point{x: p[0].x, y: p[0].y}, Point{x: p[0].x, y: p[0].y}
	for _, p := range p[1:] {
		min.x = pkgmath.MinInt(min.x, p.x)
		min.y = pkgmath.MinInt(min.y, p.y)
		max.x = pkgmath.MaxInt(max.x, p.x)
		max.y = pkgmath.MaxInt(max.y, p.y)
	}
	return &kdtree.Bounding{Min: min, Max: max}
}

func (p Points) Index(i int) kdtree.Comparable         { return p[i] }
func (p Points) Len() int                              { return len(p) }
func (p Points) Pivot(d kdtree.Dim) int                { return plane{Points: p, Dim: d}.Pivot() }
func (p Points) Slice(start, end int) kdtree.Interface { return p[start:end] }

// plane is required to help Points.
type plane struct {
	kdtree.Dim
	Points
}

func (p plane) Less(i, j int) bool {
	switch p.Dim {
	case 0:
		return p.Points[i].x < p.Points[j].x
	case 1:
		return p.Points[i].y < p.Points[j].y
	default:
		panic("illegal dimension")
	}
}
func (p plane) Pivot() int { return kdtree.Partition(p, kdtree.MedianOfMedians(p)) }
func (p plane) Slice(start, end int) kdtree.SortSlicer {
	p.Points = p.Points[start:end]
	return p
}
func (p plane) Swap(i, j int) {
	p.Points[i], p.Points[j] = p.Points[j], p.Points[i]
}
