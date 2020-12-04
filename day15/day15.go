package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dds/aoc2019/intcode"
	"github.com/dds/aoc2019/lib"
	"github.com/dds/aoc2019/lib/inputs"
	"github.com/gdamore/tcell/v2"
)

var Input = lib.InputInts(inputs.Day15(), lib.CSVParser)[0]

func main() {
	ctx, cancel := context.WithDeadline(lib.ContextWithSignals(context.Background()), time.Now().Add(120*time.Second))
	defer cancel()
	day15(ctx, Input)
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
	w, h := g.Size()
	origin := point{x: w / 2, y: h / 2}
	g.Clear()
	for p, shape := range g.cells {
		p = origin.add(p.x, -p.y)
		g.SetContent(p.x, p.y, shape, nil, tcell.StyleDefault)
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

const (
	wall    = '#'
	visited = '.'
	droid   = 'D'
)

const (
	stuck = 0
	moved = 1
	goal  = 2
)

func day15(ctx context.Context, input []int) {
	in, out := make(chan int), make(chan int)
	errs := make(chan error)
	go func() {
		if err := intcode.Code(input).Exec(ctx, in, out); err != nil {
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

	var (
		p, q                   point
		trail                  []point
		quit                   bool
		toSend                 int
		refreshRate            = time.Second / 200
		refresh                = time.NewTimer(refreshRate)
		oxygenSystem, origin   point
		distanceToOxygenSystem int
	)

	for !quit {
		if p == origin && oxygenSystem != origin {
			break
		}
		var backtracking bool
		q, toSend = g.cells.next(p)
		if toSend == NoDirection {
			backtracking = true
			g.msg = fmt.Sprintf("At %v, no place to go! trail: %v", p, trail)
			if len(trail) < 0 {
				panic(fmt.Errorf("No trail. p %v, q %v", p, q))
			}
		}
		if backtracking {
			q = trail[0]
			toSend = direction(p, q)
		}

		select {
		case <-ctx.Done():
			panic(ctx.Err())
		case err := <-errs:
			panic(err)
		case in <- toSend:
		}

		select {
		case <-ctx.Done():
			panic(ctx.Err())
		case <-userQuit:
			quit = true
		case err := <-errs:
			panic(err)
		case c, ok := <-out:
			if !ok {
				quit = true
			}
			switch c {
			case goal:
				oxygenSystem = q
				distanceToOxygenSystem = len(trail) + 1
				fallthrough
			case moved:
				g.cells[p] = visited
				g.cells[q] = droid
				if !backtracking {
					trail = append([]point{p}, trail...)
				} else {
					trail = trail[1:]
				}
				p = q
			case stuck:
				g.cells[q] = wall
			}
		}
		<-refresh.C
		refresh.Reset(refreshRate)
		g.draw()
	}

	// Part 2
	// breadth first search from oxygen tank until no more neighbors..
	refreshRate = time.Second / 50
	g.cells[oxygenSystem] = 'O'
	toExpand := []point{oxygenSystem}
	i := 0
	quit = false
	for !quit {
		for _, p := range toExpand {
			toExpand = toExpand[1:]
			for _, q := range p.neighbors() {
				if g.cells[q] == visited || g.cells[q] == droid {
					g.cells[q] = 'O'
					toExpand = append(toExpand, q)
				}
			}
		}
		if len(toExpand) == 0 {
			quit = true
		} else {
			i++
		}
		g.msg = fmt.Sprintf("At minute %v", i)
		select {
		case <-ctx.Done():
			panic(ctx.Err())
		case <-userQuit:
			quit = true
		case err := <-errs:
			panic(err)
		case <-refresh.C:
		}
		refresh.Reset(refreshRate)
		g.draw()
	}
	select {
	case <-ctx.Done():
		panic(ctx.Err())
	case <-userQuit:
	case err := <-errs:
		panic(err)
	case <-refresh.C:
	}
	g.Fini()
	fmt.Println(distanceToOxygenSystem)
	fmt.Println(i)
}
