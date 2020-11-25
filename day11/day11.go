package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/dds/aoc2019/intcode"
	"github.com/dds/aoc2019/util"
)

var Input = util.InputInts(util.Inputs[11], util.CSVParser)[0]

func main() {
	ctx := util.ContextWithSignals(context.Background())
	fmt.Println(part1(ctx, Input))
	fmt.Println(part2(ctx, Input))
}

type turn int

const (
	L turn = iota
)

func (t turn) String() string {
	if t == 0 {
		return "L"
	}
	return "R"
}

type dir int

const (
	north dir = iota
	east
	south
	west
)

func (d dir) normalize() dir {
	if d < 0 {
		d += 4
	}
	return d % 4
}

func (d dir) String() string {
	switch d := d.normalize(); d {
	case north:
		return "^"
	case east:
		return ">"
	case south:
		return "v"
	case west:
		return "<"
	}
	return ""
}

type color int

func (c color) String() string {
	if c == 0 {
		return "."
	}
	return "#"
}

func recs(ctx context.Context, i <-chan int) <-chan rec {
	recs := make(chan rec)
	go func() {
		defer close(recs)
		var r rec
		for {
			select {
			case <-ctx.Done():
				return
			case t, ok := <-i:
				if !ok {
					return
				}
				r = rec{color: color(t)}
			}
			select {
			case <-ctx.Done():
				return
			case t, ok := <-i:
				if !ok {
					return
				}
				r.turn = turn(t)
			}
			select {
			case <-ctx.Done():
				return
			case recs <- r:
			}
		}
	}()
	return recs
}

func (r rec) String() string {
	return fmt.Sprintf("{%v %v}", r.color, r.turn)
}

type rec struct {
	color color
	turn  turn
}

type point struct {
	x, y int
}

func (p *point) step(d dir) {
	switch d := d.normalize(); d {
	case north:
		p.y += 1
	case east:
		p.x += 1
	case south:
		p.y -= 1
	case west:
		p.x -= 1
	}
}

type byP []point

func (a byP) Len() int      { return len(a) }
func (a byP) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byP) Less(i, j int) bool {
	if a[i].y < a[j].y {
		return true
	}
	return a[i].x < a[j].x
}

type board map[point]color

func (b board) String() string {
	if len(b) == 0 {
		return ""
	}
	var minP, maxP point
	points := byP{}
	for p := range b {
		if p.x < minP.x {
			minP.x = p.x
		}
		if p.y < minP.y {
			minP.y = p.y
		}
		if p.x > maxP.x {
			maxP.x = p.x
		}
		if p.y > maxP.y {
			maxP.y = p.y
		}
		points = append(points, p)
	}
	sort.Sort(points)

	w := maxP.x - minP.x
	h := maxP.y - minP.y
	d := make([][]byte, h+1)
	for y := 0; y < h+1; y++ {
		row := make([]byte, w+1)
		for x := 0; x < w+1; x++ {
			row[x] += '.'
		}
		d[y] = row
	}

	T := func(p point) point {
		p.x -= minP.x
		p.y -= minP.y
		return p
	}
	for _, p := range points {
		if b[p] != 1 {
			continue
		}
		q := T(p)
		d[q.y][q.x] = '#'
	}
	s := "\n\n"
	for y := h; y >= 0; y-- {
		s += string(d[y]) + "\n"
	}
	return s
}

func part1(ctx context.Context, input []int) (r int) {
	out := make(chan int)
	in := make(chan int, 1)
	in <- 0
	go func() {
		if err := intcode.Code(input).Exec(ctx, in, out); err != nil {
			panic(err)
		}
	}()

	recs := recs(ctx, out)

	var m = make(board)
	var p point
	var d dir
	for {
		// Process output
		select {
		case <-ctx.Done():
			return
		case t, ok := <-recs:
			if !ok {
				return len(m)
			}
			m[p] = t.color
			if t.turn == L {
				d--
			} else {
				d++
			}
			d = d.normalize()
			p.step(d)
		}
		// Send input
		select {
		case <-ctx.Done():
			return len(m)
		case in <- int(m[p]):
		}
	}
}

func part2(ctx context.Context, input []int) (r int) {
	out := make(chan int)
	in := make(chan int, 1)
	in <- 1
	go func() {
		if err := intcode.Code(input).Exec(ctx, in, out); err != nil {
			panic(err)
		}
	}()

	recs := recs(ctx, out)

	var m = make(board)
	var p point
	var d dir
	for {
		// Process output
		select {
		case <-ctx.Done():
			return
		case t, ok := <-recs:
			if !ok {
				return len(m)
			}
			m[p] = t.color
			if t.turn == L {
				d--
			} else {
				d++
			}
			d = d.normalize()
			p.step(d)
		}
		fmt.Println(m)
		// Send input
		select {
		case <-ctx.Done():
			return len(m)
		case in <- int(m[p]):
		}
	}
}
