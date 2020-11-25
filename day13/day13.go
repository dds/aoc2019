package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dds/aoc2019/intcode"
	"github.com/dds/aoc2019/lib"
)

var Input = lib.InputInts(lib.Inputs[13], lib.CSVParser)[0]

func main() {
	ctx := lib.ContextWithSignals(context.Background())
	fmt.Println(part1(ctx, Input))
	fmt.Println(part2(ctx, Input))
}

type shape int

type rec struct {
	x, y int
	id   shape
}

func (r rec) String() string {
	return fmt.Sprintf("(%d, %d): %q", r.x, r.y, r.id)
}

const (
	empty shape = iota
	wall
	block
	paddle
	ball
)

var shapes = map[shape]string{
	empty:  " ",
	wall:   "=",
	block:  "#",
	paddle: "|",
	ball:   "o",
}

func (s shape) String() string {
	return shapes[s]
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
				r.id = shape(t)
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
		if r.id != block {
			continue
		}
		i++
	}
	return i
}

func part2(ctx context.Context, input []int) (r int) {
	in := make(chan int)
	out := make(chan int)
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(3*time.Second))
	defer cancel()
	input[0] = 2
	go func() {
		if err := intcode.Code(input).Exec(ctx, in, out); err != nil {
			panic(err)
		}
	}()
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			return
	// 		case in <- 1:
	// 		}
	// 	}
	// }()
	for i := range recs(ctx, out) {
		fmt.Println(i)
	}
	return 0
}
