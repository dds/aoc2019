package main

import (
	"fmt"
	"math"

	"github.com/dds/aoc2020/util"
)

var Input = util.Inputs[10]

func main() {
	fmt.Println(part1(Input))
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
