package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dds/aoc2019/intcode"
	"github.com/dds/aoc2019/lib"
	"github.com/gdamore/tcell"
)

var Input = lib.InputInts(lib.Inputs[13], lib.CSVParser)[0]

func main() {
	ctx := lib.ContextWithSignals(context.Background())
	fmt.Println(part1(ctx, Input))
	fmt.Println(part2(ctx, Input))
}

type shape int

type rec struct {
	x, y  int
	shape shape
}

func (r rec) String() string {
	return fmt.Sprintf("(%d, %d): %v", r.x, r.y, r.shape)
}

const (
	empty shape = iota
	wall
	block
	paddle
	ball
)

var shapes = map[shape]rune{
	empty:  ' ',
	wall:   '|',
	block:  '#',
	paddle: '-',
	ball:   'o',
}

func (s shape) String() string {
	return string(shapes[s])
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
				r = rec{x: t}
			}
			select {
			case <-ctx.Done():
				return
			case t, ok := <-i:
				if !ok {
					return
				}
				r.y = t
			}
			select {
			case <-ctx.Done():
				return
			case t, ok := <-i:
				if !ok {
					return
				}
				r.shape = shape(t)
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

func part1(ctx context.Context, input []int) (r int) {
	in := make(chan int)
	out := make(chan int)
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(4*time.Second))
	defer cancel()
	go func() {
		if err := intcode.Code(input).Exec(ctx, in, out); err != nil {
			panic(err)
		}
	}()
	var i = 0
	for r := range recs(ctx, out) {
		if r.shape != block {
			continue
		}
		i++
	}
	return i
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type point struct {
	x, y int
}

func part2(ctx context.Context, input []int) (rc int) {
	input[0] = 2
	// Boot intcode computer.
	in := make(chan int)
	out := make(chan int)
	errs := make(chan error)
	go func() {
		if err := intcode.Code(input).Exec(ctx, in, out); err != nil {
			errs <- err
		}
	}()
	recs := recs(ctx, out)

	var err error

	// Setup screen.
	// encoding.Register()
	// cb := &tcell.CellBuffer{}
	// fmt.Println(cb)

	var sc tcell.Screen
	sc, err = tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := sc.Init(); err != nil {
		panic(err)
	}
	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	sc.SetStyle(style)
	// sc.Clear()

	userQuit := make(chan struct{})
	go func() {
		for {
			switch ev := sc.PollEvent().(type) {
			case *tcell.EventResize:
				sc.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					userQuit <- struct{}{}
				}
			}
		}
	}()

	// Start game loop.
	var (
		quit bool
		p    point
		w, h int
	)
	refreshRate := time.Second / 360
	refresh := time.NewTimer(refreshRate)
	const notpending = -2
	var pending = notpending
	var pendingC chan int
	for !quit {
		if pending != notpending {
			// If we have a pending send, set its channel
			pendingC = in
		}
		// Setup pending input to game.
		select {
		case err = <-errs:
			quit = true
		case <-ctx.Done():
			// Quit if canceled.
			quit = true
		case <-userQuit:
			quit = true
		case pendingC <- pending:
			// Send pending input to game.
			pendingC = nil
			pending = notpending
		case u, ok := <-recs:
			if !ok {
				quit = true
			}
			w = max(w, u.x)
			h = max(h, u.y)
			switch u.shape {
			case paddle:
				p.x = u.x
				p.y = u.y
			case ball:
				if p.x == u.x {
					pending = 0
				} else if p.x < u.x {
					pending = 1
				} else {
					pending = -1
				}
			}
			// Update screen based on output.
			if u.x == -1 && u.y == 0 {
				u.y = h + 1
				u.x = (w / 2) - 2
				rc = int(u.shape)
			}
			sc.SetContent(u.x, u.y, shapes[u.shape], nil, tcell.StyleDefault)
			for i, c := range fmt.Sprint(rc) {
				sc.SetContent(w/2+i, h+1, c, nil, tcell.StyleDefault)
			}
			// sc.SetContent(u.x, u.y, shapes[u.shape], nil, 0)
		}
		<-refresh.C
		// Redraw on timer elapse.
		refresh.Reset(refreshRate)
		// Update score
		// Draw board
		sc.Show()
	}
	sc.Fini()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	return
}
