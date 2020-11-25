package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/dds/aoc2019/lib"
)

var Input = lib.Inputs[10]

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}

const (
	space = '.'
)

func Parse(input string) [][]byte {
	b := [][]byte{}
	row := []byte{}
	for _, c := range input {
		switch c := byte(c); c {
		case '\n':
			if len(row) > 0 {
				b = append(b, row)
				row = []byte{}
			}
			continue
		case space:
			fallthrough
		default:
			row = append(row, c)
		}
	}
	return b
}

type point struct {
	x, y int
}

func part1(input string) (int, point) {
	m := map[point]int{}
	points := Parse(input)

	fmt.Println("W", len(points[0]), "H", len(points))
	for y, row := range points {
		for x, b := range row {
			if b == space {
				fmt.Print(string(space))
				continue
			}
			fmt.Print("#")
			m[point{x, y}] = 0
		}
		fmt.Printf("\n")
	}

	var bestPoint point
	var bestScore int
	for p := range m {
		slopes := map[float64]int{}
		for q := range m {
			if p == q {
				continue
			}
			slope := math.Atan2(float64(q.y-p.y), float64(q.x-p.x))
			if slopes[slope] == 0 {
				slopes[slope] = 1
			}
		}
		m[p] = len(slopes)
	}
	for y, row := range points {
		for x := range row {
			var score int
			var ok bool
			p := point{x, y}
			if score, ok = m[p]; !ok {
				continue
			}
			if score > bestScore {
				bestScore = score
				bestPoint = p
			}
		}
	}
	return bestScore, bestPoint
}

func Field(input [][]byte) (r []point) {
	for y, row := range input {
		for x, b := range row {
			if b == space {
				continue
			}
			p := point{x, y}
			r = append(r, p)
		}
	}
	return
}

func part2(input string) (int, point) {
	m := map[point]int{}
	points := Parse(input)

	for y, row := range points {
		for x, b := range row {
			if b == space {
				continue
			}
			m[point{x, y}] = 0
		}
	}

	var basePoint point
	var bestScore int
	for p := range m {
		slopes := map[float64]int{}
		for q := range m {
			if p == q {
				continue
			}
			slope := math.Atan2(float64(q.y-p.y), float64(q.x-p.x))
			if slopes[slope] == 0 {
				slopes[slope] = 1
			}
		}
		m[p] = len(slopes)
	}
	for y, row := range points {
		for x := range row {
			var score int
			var ok bool
			p := point{x, y}
			if score, ok = m[p]; !ok {
				continue
			}
			if score > bestScore {
				bestScore = score
				basePoint = p
			}
		}
	}
	field := Field(points)

	targets := vaporize(basePoint, field)
	if len(targets) >= 200 {
		return 100*targets[199].x + targets[199].y, targets[199].point
	}
	return 0, basePoint
}

type target struct {
	point
	r angle
	d float64
}

type byR []target

func (a byR) Len() int           { return len(a) }
func (a byR) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byR) Less(i, j int) bool { return a[i].r < a[j].r }

type angle float64

// Normalized returns an equivalent angle in (0, 2π].
func (a angle) normalized() angle {
	rad := math.Remainder(float64(a), 2*math.Pi)
	if rad <= -math.Pi {
		rad = math.Pi
	}
	rad += math.Pi / 2.0
	if rad < 0 {
		rad += 2 * math.Pi
	}
	return angle(rad)
}

func vaporize(base point, field []point) []target {
	m := map[angle]target{}
	targets := byR{}
	for _, p := range field {
		if p == base {
			continue
		}
		dy := float64(p.y - base.y)
		dx := float64(p.x - base.x)
		d := math.Hypot(dx, dy)
		r := angle(math.Atan2(dy, dx)).normalized()
		t := target{p, r, d}
		if m[r].d != 0 && d >= m[r].d {
			continue
		}
		m[r] = t
	}
	for _, t := range m {
		targets = append(targets, t)
	}
	sort.Sort(targets)
	return targets
}
