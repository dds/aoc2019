package main

import (
	"context"
	"fmt"

	"github.com/dds/aoc2019/intcode"
	"github.com/dds/aoc2019/util"
)

var Input = util.InputInts(util.Inputs[11], util.CSVParser)[0]

func main() {
	ctx := util.ContextWithSignals(context.Background())
	fmt.Println(part1(ctx, Input))
	// fmt.Println(part2(ctx, Input))
}

type dir int

const (
	Up dir = iota
	Right
	Down
	Left
)

func (d dir) normalize() dir {
	return d % 4
}

func (d dir) String() string {
	switch d.normalize() {
	case Up:
		return "^"
	case Right:
		return ">"
	case Down:
		return "v"
	case Left:
		return "<"
	}
	return ""
}

type color int

const (
	Black color = iota
	White
)

func (c color) String() string {
	if c == 0 {
		return "."
	}
	return "#"
}

type rec struct {
	color color
	dir   dir
}

func (r rec) String() string {
	return fmt.Sprintf("{%v %v}", r.color, r.dir)
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
				r.dir = dir(t)
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

func part1(ctx context.Context, input []int) (r interface{}) {
	out := make(chan int)
	in := make(chan int)
	go intcode.Code(input).Exec(ctx, in, out)

	recs := recs(ctx, out)

	go func() {
		for i := 0; i < 10; i++ {
			in <- 0
		}
		close(in)
	}()

	s := []rec{}
	for i := range recs {
		s = append(s, i)
	}
	return fmt.Sprint(s)
}

func part2(ctx context.Context, input []int) (r interface{}) {
	out := make(chan int)
	in := make(chan int)
	go intcode.Code(input).Exec(ctx, in, out)
	in <- 2
	s := []int{}
	for i := range out {
		s = append(s, i)
	}
	return fmt.Sprint(s)
}
