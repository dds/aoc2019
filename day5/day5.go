package main

import (
	"context"
	"fmt"

	"github.com/dds/aoc2019/intcode"
	"github.com/dds/aoc2019/lib"
)

var Input = lib.InputInts(lib.Inputs[5], lib.CSVParser)[0]

func main() {
	ctx := lib.ContextWithSignals(context.Background())
	fmt.Println(part1(ctx, Input))
	fmt.Println(part2(ctx, Input))
}

func part1(ctx context.Context, input []int) string {
	out := make(chan int)
	in := make(chan int)
	go intcode.Code(input).Exec(ctx, in, out)
	in <- 1
	s := []int{}
	for i := range out {
		s = append(s, i)
	}
	return fmt.Sprint(s)
}

func part2(ctx context.Context, input []int) string {
	out := make(chan int)
	in := make(chan int)
	go intcode.Code(input).Exec(ctx, in, out)
	in <- 5
	s := []int{}
	for i := range out {
		s = append(s, i)
	}
	return fmt.Sprint(s)
}
