package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dds/aoc2019/intcode"
	"github.com/dds/aoc2019/lib"
	"github.com/dds/aoc2019/lib/inputs"
	"github.com/gdamore/tcell"
)

var Input = lib.InputInts(inputs.Day17(), lib.NumberParser)[0]

func main() {
	ctx, cancel := context.WithDeadline(lib.ContextWithSignals(context.Background()), time.Now().Add(30*time.Second))
	defer cancel()
	fmt.Println(part1(ctx, Input))
	fmt.Println(part2(ctx, Input))
}

func part1(ctx context.Context, input []int) (rc int) {
	in, out := make(chan int), make(chan int)
	code := intcode.Code(input)
	errs := make(chan error)
	go func() {
		if err := code.Exec(ctx, in, out); err != nil {
			errs <- err
		}
	}()

	userQuit := make(chan struct{})
	g := &grid{cells: make(cells)}
	g.init()
	g.draw()
	go func() {
		for {
			switch ev := g.PollEvent().(type) {
			case *tcell.EventResize:
				g.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					userQuit <- struct{}{}
				}
			}
		}
	}()
	defer func() {
		g.Fini()
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	w, h := 0, 0
	x, y := 0, 0
	for c := range out {
		switch c := rune(c); c {
		case '\n':
			x = 0
			y++
		default:
			g.cells[point{x, y}] = c
			x++
		}
		w, h = lib.Max(w, x), lib.Max(h, y)
	}
cells:
	for p, _ := range g.cells {
		count := 0
		for _, neighbor := range p.neighbors() {
			if g.cells[p] != '#' || g.cells[neighbor] != '#' {
				continue cells
			}
			count++
		}
		if count < 4 {
			continue cells
		}
		g.cells[p] = 'O'
		rc += p.x * p.y
	}
	g.draw()
	select {
	case <-ctx.Done():
		panic(ctx.Err())
	case <-userQuit:
	}
	g.Fini()
	fmt.Println("WxH: ", w, "x", h)
	return
}

func part2(ctx context.Context, input []int) (rc int) {
	in, out := make(chan int), make(chan int)
	code := intcode.Code(input)
	code[0] = 2
	errs := make(chan error)
	go func() {
		if err := code.Exec(ctx, in, out); err != nil {
			errs <- err
		}
	}()

	go func() {
		for {
			for _, c := range fmt.Sprintf("A,A,B,C,B,C,B,C\n10,L,8,R,6\n") {
				select {
				case <-ctx.Done():
				case in <- int(c):
				}
			}
		}
	}()
	userQuit := make(chan struct{})
	g := &grid{cells: make(cells)}
	g.init()
	g.draw()
	go func() {
		for {
			switch ev := g.PollEvent().(type) {
			case *tcell.EventResize:
				g.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					userQuit <- struct{}{}
				}
			}
		}
	}()
	defer func() {
		g.Fini()
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	g.draw()
	select {
	case <-ctx.Done():
		panic(ctx.Err())
	case <-userQuit:
	}
	g.Fini()
	return
	return
}

type point struct {
	x, y int
}

func (p point) add(x, y int) point {
	p.x += x
	p.y += y
	return p
}

type cells map[point]rune

var directions = []point{
	point{0, 1},
	point{0, -1},
	point{-1, 0},
	point{1, 0},
}

const NoDirection = -1

func direction(p, q point) int {
	d := q.add(-p.x, -p.y)
	if d.x != 0 {
		d.x = lib.Sign(d.x)
	}
	if d.y != 0 {
		d.y = lib.Sign(d.y)
	}
	for i, v := range directions {
		if v == d {
			return i + 1
		}
	}
	return NoDirection
}

func (p point) neighbors() (r []point) {
	for _, q := range directions {
		r = append(r, p.add(q.x, q.y))
	}
	return
}

func (c cells) next(p point) (point, int) {
	// Find the first unexplored direction.
	for i, q := range p.neighbors() {
		if _, ok := c[q]; !ok {
			return q, i + 1
		}
	}
	return point{}, NoDirection
}

type grid struct {
	tcell.Screen
	cells
	scene int
	msg   string
}

func (g *grid) init() {
	sc, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := sc.Init(); err != nil {
		panic(err)
	}
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	sc.SetStyle(style)

	g.Screen = sc
	g.cells = make(cells)
}

func (g *grid) draw() {
	g.Clear()
	for p, shape := range g.cells {
		g.SetContent(p.x, p.y+4, shape, nil, tcell.StyleDefault)
	}
	g.scene++
	for i, c := range fmt.Sprintf("Scene: %v", g.scene) {
		g.SetContent(i, 0, c, nil, tcell.StyleDefault)
	}
	var i, line int = 0, 1
	for _, ch := range g.msg {
		if ch == '\n' {
			line++
			i = 0
			continue
		}
		g.SetContent(i, line, ch, nil, tcell.StyleDefault)
		i++
	}
	g.Show()
}
