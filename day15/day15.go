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
	ctx, cancel := context.WithDeadline(lib.ContextWithSignals(context.Background()), time.Now().Add(120*time.Second))
	defer cancel()
	fmt.Println(part1(ctx, Input))
	fmt.Println(part2(ctx, Input))
}

type point struct {
	x, y int
}

func (p point) add(x, y int) point {
	p.x += x
	p.y += y
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
		fmt.Println("... and we're back")
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	var (
		p, q        point
		w, h        int
		trail       []point
		quit        bool
		toSend      int
		refreshRate = time.Second / 1000
		refresh     = time.NewTimer(refreshRate)
	)

	g.cells[p] = 'O'
	var oxygenSystem point
	for !quit {
		var backtracking bool
		q, toSend = g.cells.next(p)
		if toSend == -1 {
			backtracking = true
			g.msg = fmt.Sprintf("At %v, no place to go! trail: %v", p, trail)
			if len(trail) < 0 {
				panic(fmt.Errorf("No trail. p %v, q %v", p, q))
			}
			q = trail[0]
			toSend = direction(p, q)
			g.msg += fmt.Sprintf("\nFinding direction: p %v, q %v, toSend %v", p, q, toSend)
		} else {
			g.msg = fmt.Sprintf("At %v, trying %v, trail: %v", p, q, trail)
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
			g.msg += fmt.Sprintf("\nReceived: %v, %v", c, ok)
			switch c {
			case goal:
				g.cells[q] = destination
				g.cells[p] = droid
				quit = true
				oxygenSystem = q
			case moved:
				if !(p.x == 0 && p.y == 0) {
					g.cells[p] = visited
				}
				g.cells[q] = droid
				if !backtracking {
					trail = append([]point{p}, trail...)
				} else {
					trail = trail[1:]
				}
				g.msg += fmt.Sprintf("\nUpdating point from %v => %v", p, q)
				p = q
				w = lib.Max(w, 2*lib.Sign(p.x)*p.x)
				h = lib.Max(w, 2*lib.Sign(p.y)*p.y)
			case stuck:
				g.cells[q] = wall
			}
		}
		<-refresh.C
		refresh.Reset(refreshRate)
		g.draw()
	}
	g.msg += fmt.Sprintf("\nOxygen system is at %v\nGrid is %vx%v. Press escape to quit.", oxygenSystem, w, h)
	g.draw()
	select {
	case <-ctx.Done():
	case <-userQuit:
	case err := <-errs:
		panic(err)
	}
	fmt.Printf("Found oxygen system at %v, finding optimal path...\n", oxygenSystem)
	optimalPath, err := g.astar(point{0, 0}, oxygenSystem)
	if err != nil {
		panic(err)
	}
	refreshRate = time.Second / 2
	for _, p := range optimalPath {
		g.cells[p] = '+'
		g.draw()
		select {
		case <-ctx.Done():
		case <-userQuit:
		case err := <-errs:
			panic(err)
		case <-refresh.C:
			refresh.Reset(refreshRate)
		}
	}
	select {
	case <-ctx.Done():
	case <-userQuit:
	case err := <-errs:
		panic(err)
	case <-refresh.C:
		refresh.Reset(refreshRate)
	}
	fmt.Println(len(optimalPath))
	return
}

var ErrNoPath = fmt.Errorf("no path")

func (m cells) astar(start point, goal point) (trail []point, err error) {
	// type visit struct {
	// 	parent point
	// 	distanceFromStart
	// }
	// visited := map[point]visit{} // map visited node to previous nodes

	// pq := pq{&node{point: start}}
	// heap.Init(&pq)
	// p := start

	// for pq.Len() > 0 {
	// 	u := heap.Pop(&pq).(*node)
	// 	visited[u.point] = visit{parent: p, distanceFromStart: visit[p].distanceFromStart + 1}
	// 	p = u.point

	// 	if u.point == goal {
	// 		for {
	// 			trail = append([]point{u.point}, trail...)
	// 			if p, ok := visited[parent]; ok {
	// 				u.point = p.parent
	// 				continue
	// 			}
	// 			return
	// 		}
	// 	}

	// 	// time.Sleep(time.Second / 2)
	// 	for _, q := range u.point.neighbors() {
	// 		if _, ok := visited[q]; ok {
	// 			continue
	// 		}
	// 		movementCost := 1
	// 		if c, ok := m[q]; c == wall {
	// 			movementCost += 1 << 20
	// 		} else if !ok {
	// 			movementCost += 1 << 10
	// 		}
	// 		cost := len(trail) + movementCost
	// 		in, pos := pq.Find(q)
	// 		if in != nil && cost < in.cost {
	// 			heap.Remove(&pq, pos) // alternative path is better
	// 		}
	// 		var seen bool
	// 		_, seen = visited[q]
	// 		// If we have passed this point, restore it if we found a better way to get there.
	// 		if in != nil && seen && cost < in.cost {
	// 			delete(visited, q)
	// 		}
	// 		if in == nil && !seen {
	// 			in = &node{point: q, cost: cost + q.distance(goal)}
	// 			heap.Push(&pq, in)
	// 		}
	// 	}
	// }
	// err = ErrNoPath
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

func (pq pq) Find(p point) (*node, int) {
	for i, q := range pq {
		if p == q.point {
			return q, i
		}
	}
	return nil, -1
}

func (pq pq) Len() int { return len(pq) }

func (pq pq) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq pq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *pq) Push(x interface{}) {
	n := len(*pq)
	node := x.(*node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *pq) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *pq) update(item *node, score int, priority int) {
	item.cost = score
	heap.Fix(pq, item.index)
}
