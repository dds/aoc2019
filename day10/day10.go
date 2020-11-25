package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/dds/aoc2019/util"
)

var Input = util.Inputs[10]

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

func part2(input string) (int, point) {
	m := map[point]int{}
	points := Parse(input)

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
	field := []point{}
	for y, row := range points {
		for x := range row {
			var score int
			var ok bool
			p := point{x, y}
			field = append(field, p)
			if score, ok = m[p]; !ok {
				continue
			}
			if score > bestScore {
				bestScore = score
				bestPoint = p
			}
		}
	}

	from := bestPoint
	targets := vaporize(from, field)
	fmt.Println("targets", targets)
	n := 0
	for {
		for i := 0; i < len(field); i++ {
			p := field[i]
			for _, t := range targets {
				if t.point != p {
					continue
				}
				// Fire when ready
				fmt.Println("DESTROYING", p)
				field[i] = field[len(field)-1]
				field[len(field)-1] = point{}
				field = field[:len(field)-1]
				n++
				if n == 200 {
					return 100*t.x + t.y, t.point
				}
			}
		}
		targets = vaporize(from, field)
		if len(targets) == 0 {
			break
		}
	}
	return 0, from
}

type target struct {
	point
	r, d float64
}

type byR []target

func (a byR) Len() int           { return len(a) }
func (a byR) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byR) Less(i, j int) bool { return a[i].r < a[j].r }

func vaporize(from point, field []point) []target {
	m := map[float64]target{}
	targets := byR{}
	for _, p := range field {
		dy := float64(from.y - p.y)
		dx := float64(from.x - p.x)
		r := math.Atan2(dy, dx)
		d := dx*dx + dy*dy
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
