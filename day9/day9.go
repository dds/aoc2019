package main

import (
	"context"
	"fmt"

	"github.com/dds/aoc2020/intcode"
	"github.com/dds/aoc2020/util"
)

var Input = util.InputInts(util.Inputs[9], util.CSVParser)[0]

func main() {
	ctx := util.ContextWithSignals(context.Background())
	fmt.Println(part1(ctx, Input))
	fmt.Println(part2(ctx, Input))
}

func part1(ctx context.Context, input []int) (r interface{}) {
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

func part2(ctx context.Context, input []int) (r interface{}) {
	return input
}
