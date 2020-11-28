package main

import (
	"container/heap"
	"context"
	"fmt"
	"time"

	"github.com/dds/aoc2019/intcode"
	"github.com/dds/aoc2019/lib"
	"github.com/dds/aoc2019/lib/inputs"
	"github.com/gdamore/tcell"
)

var Input = lib.InputInts(inputs.Day15(), lib.CSVParser)[0]

func main() {
	ctx, cancel := context.WithDeadline(lib.ContextWithSignals(context.Background()), time.Now().Add(10*time.Second))
	defer cancel()
	fmt.Println(part1(ctx, Input))
	fmt.Println(part2(ctx, Input))
}

type point struct {
	x, y int
}

func (p point) add(q point) point {
	p.x = p.x + q.x
	p.y = p.y + q.y
	return p
}

// Estimate the distance from the manhattan distance.
func (p point) distance(q point) int {
	var p1, p2 int
	if q.x > p.x {
		p1 = q.x - p.x
	} else {
		p1 = p.x - q.x
	}
	if q.y > p.y {
		p2 = q.y - p.y
	} else {
		p2 = p.y - q.y
	}
	return p1 + p2
}

type cells map[point]rune

type grid struct {
	tcell.Screen
	cells
	scene int
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
		p = origin.add(p)
		g.SetContent(p.x, p.y, shape, nil, tcell.StyleDefault)
	}
	g.scene++
	for i, c := range fmt.Sprintf("Scene: %v", g.scene) {
		g.SetContent(i, 0, c, nil, tcell.StyleDefault)
	}
	g.Show()
}

const (
	empty       = ' '
	wall        = '#'
	visited     = '.'
	droid       = 'D'
	destination = 'G'
)

// input values
const (
	north = 1
	south = 2
	west  = 3
	east  = 4
)

// output values
const (
	stuck = 0
	moved = 1
	goal  = 2
)

func part1(ctx context.Context, input []int) (rc int) {
	in, out := make(chan int), make(chan int)
	errs := make(chan error)
	go func() {
		if err := intcode.Code(input).Exec(ctx, in, out); err != nil {
			select {
			case <-ctx.Done():
				return
			case errs <- err:
			}
		}
	}()

	g := &grid{}
	g.init()
	g.draw()

	defer func() {
		g.Fini()
		fmt.Println("... and we're back")
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	userQuit := make(chan struct{})
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

	var (
		p      point
		quit   bool
		toSend int
	)

	const noInput = -1

	for !quit {
		var inputC chan int
		toSend = east
		if toSend != noInput {
			inputC = in
		}

		q := p.add(point{1, 0})

		select {
		case <-ctx.Done():
			panic(ctx.Err())
		case <-userQuit:
			quit = true
		case err := <-errs:
			panic(err)
		case inputC <- toSend:
			toSend = noInput
		case c := <-out:
			switch c {
			case goal:
				g.cells[q] = destination
				quit = true
			case moved:
				g.cells[p] = visited
				p = q
			case stuck:
				g.cells[q] = wall
			}
		}
		//  - paint the screen
		g.cells[p] = droid
		g.draw()
	}
	//  - a_star on the location
	return
}

func part2(ctx context.Context, input []int) (rc int) {
	return 0
}

type node struct {
	point
	cost int
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type pq []*node

func (pq pq) Len() int { return len(pq) }

func (pq pq) Less(i, j int) bool {
	return pq[i].cost > pq[j].cost
}

func (pq pq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pq) Push(n interface{}) {
	node := n.(node)
	node.index = len(*pq)
	*pq = append(*pq, &node)
}

func (pq *pq) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return *item
}

func (pq *pq) update(item *node, score int, priority int) {
	item.cost = score
	heap.Fix(pq, item.index)
}
