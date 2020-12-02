package main

import (
	"context"
	"fmt"
	"image"
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

var W, H = 60, 40

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
	g := intcode.NewGrid()
	g.Draw()
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

	x, y := 0, 0
	for c := range out {
		switch c := rune(c); c {
		case '\n':
			x = 0
			y++
		default:
			g.Cells()[image.Point{x, y}] = intcode.Cell{c, tcell.StyleDefault}
			x++
		}
	}
cells:
	for p, _ := range g.Cells() {
		count := 0
		for _, neighbor := range intcode.Neighbors(p) {
			if g.Cells()[p].Rune != '#' || g.Cells()[neighbor].Rune != '#' {
				continue cells
			}
			count++
		}
		if count < 4 {
			continue cells
		}
		g.Cells()[p] = intcode.Cell{'O', tcell.StyleDefault}
		rc += p.X * p.Y
	}
	g.Draw()
	select {
	case <-ctx.Done():
		panic(ctx.Err())
	case <-userQuit:
	}
	g.Fini()
	return
}

func part2(ctx context.Context, input []int) (rc int) {
	return
}
