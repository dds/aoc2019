package main

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/dds/aoc2020/intcode"
	"github.com/dds/aoc2020/util"
)

var Input = util.InputInts(util.Inputs[7], util.CSVParser)[0]

func main() {
	ctx := util.ContextWithSignals(context.Background())
	fmt.Println(part1(ctx, Input))
	fmt.Println(part2(ctx, Input))
}

type amp struct {
	code    intcode.Code
	in, out chan int
	phase   int
}

func (a *amp) Run(ctx context.Context) error {
	return a.code.Exec(ctx, a.in, a.out)
}

const n = 5

var ErrInvalid = errors.New("invalid")

func parsePhase(phase int) ([]int, error) {
	m := map[int]bool{}
	r := make([]int, n)

	p := phase
	for i := n - 1; i >= 0; i-- {
		x := p % 10
		if x > 4 {
			return nil, fmt.Errorf("%w: %v: value too large %v", ErrInvalid, phase, x)
		}
		if m[x] {
			return nil, fmt.Errorf("%w: %v: repeat element %v", ErrInvalid, phase, x)
		}
		m[x] = true
		p = p / 10
		r[i] = x
	}
	return r, nil
}

func Part1(ctx context.Context, code intcode.Code, phases []int) (r int) {
	amps := make([]*amp, len(phases))
	for i := range phases {
		c := make([]int, len(code))
		copy(c, code)
		amps[i] = &amp{
			in:    make(chan int),
			out:   make(chan int),
			code:  intcode.Code(c),
			phase: phases[i],
		}
		go func(a *amp) {
			if err := a.Run(ctx); err != nil {
				panic(err)
			}
		}(amps[i])
	}

	for _, a := range amps {
		select {
		case <-ctx.Done():
			return
		case a.in <- a.phase:
		}
		select {
		case <-ctx.Done():
			return
		case a.in <- r:
		}
		select {
		case <-ctx.Done():
			return
		case r = <-a.out:
		}
	}

	return
}

func part1(ctx context.Context, input []int) (r int) {
	c := intcode.Code(input)

	for i := 0; i < 60000; i++ {
		p, err := parsePhase(i)
		if err != nil {
			continue
		}
		a := Part1(ctx, c, p)
		if a > r {
			r = a
		}
	}
	return
}

func parsePhase2(phase int) ([]int, error) {
	m := map[int]bool{}
	r := make([]int, n)

	p := phase
	for i := n - 1; i >= 0; i-- {
		x := p % 10
		if x < 5 {
			return nil, fmt.Errorf("%w: %v: value too small %v", ErrInvalid, phase, x)
		}
		if m[x] {
			return nil, fmt.Errorf("%w: %v: repeat element %v", ErrInvalid, phase, x)
		}
		m[x] = true
		p = p / 10
		r[i] = x
	}
	return r, nil
}

func Part2(ctx context.Context, code intcode.Code, phases []int) (r int) {
	amps := make([]*amp, len(phases))
	var wg sync.WaitGroup
	for i := range phases {
		c := make([]int, len(code))
		copy(c, code)
		amps[i] = &amp{
			in:    make(chan int, 1),
			out:   make(chan int),
			code:  intcode.Code(c),
			phase: phases[i],
		}
		go func(a *amp) {
			defer wg.Done()
			if err := a.Run(ctx); err != nil {
				panic(err)
			}
		}(amps[i])
		wg.Add(1)
	}

	for i := range amps {
		amps[i].in <- phases[i]
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	for i := 0; ; i++ {
		a := amps[i%len(amps)]
		select {
		case <-ctx.Done():
			panic(ctx.Err())
		case <-done:
			return
		case a.in <- r:
		}
		select {
		case <-ctx.Done():
			panic(ctx.Err())
		case <-done:
			return
		case t, ok := <-a.out:
			if !ok {
				continue
			}
			r = t
		}
	}
}

func part2(ctx context.Context, input []int) (r int) {
	c := intcode.Code(input)

	for i := 55555; i < 100000; i++ {
		p, err := parsePhase2(i)
		if err != nil {
			continue
		}
		a := Part2(ctx, c, p)
		if a > r {
			r = a
		}
	}
	return
}
