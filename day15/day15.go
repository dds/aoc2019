package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dds/aoc2019/intcode"
	"github.com/dds/aoc2019/lib"
	"github.com/dds/aoc2019/lib/inputs"
)

var Input = lib.InputInts(inputs.Day15(), lib.CSVParser)[0]

func main() {
	ctx := lib.ContextWithSignals(context.Background())
	fmt.Println(part1(ctx, Input))
	fmt.Println(part2(ctx, Input))
}

type shape int

type point struct {
	x, y int
}
type rec struct {
	point point
	shape shape
}

func (r rec) String() string {
	return fmt.Sprintf("{%v, %v}", r.point, r.shape)
}

const (
	empty shape = iota
	wall
	visited
	droid
)

var shapes = map[shape]rune{
	empty:   ' ',
	wall:    '#',
	visited: '.',
	droid:   'D',
}

func (s shape) String() string {
	return string(shapes[s])
}

const (
	north = 1
	south = 2
	west  = 3
	east  = 4
)

const (
	stuck = 0
	moved = 1
	done  = 2
)

func part1(ctx context.Context, input []int) (rc int) {
	in := make(chan int)
	out := make(chan int)
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()
	go func() {
		if err := intcode.Code(input).Exec(ctx, in, out); err != nil {
			panic(err)
		}
	}()

	go func() {
		for i := 0; ; i++ {
			var step = i % 4
			switch step {
			// convert 0 1 2 3 -> 1 4 2 3
			case 1:
				step = 3
			case 2:
				step = 1
			case 3:
				step = 2
			}
			step++
			fmt.Println("moving: ", step)
			in <- step
		}
	}()

	for i := range out {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	return 0
}

func part2(ctx context.Context, input []int) (rc int) {
	return 0
}
