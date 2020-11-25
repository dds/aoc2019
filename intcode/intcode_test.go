package intcode_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/dds/aoc2019/intcode"
	"github.com/stretchr/testify/require"
)

func TestIO(t *testing.T) {
	type input struct {
		code, in []int
	}
	type test struct {
		input  input
		expect []int
	}

	tests := map[string]test{
		"1st": test{
			input: input{
				code: []int{3, 0, 4, 0, 99},
				in:   []int{1},
			},
			expect: []int{1},
		},
		"2nd": test{
			input: input{
				code: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
				in:   []int{8},
			},
			expect: []int{1},
		},
		"3rd": test{
			input: input{
				code: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
				in:   []int{9},
			},
			expect: []int{0},
		},
		"4th": test{
			input: input{
				code: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
				in:   []int{7},
			},
			expect: []int{1},
		},
		"5th": test{
			input: input{
				code: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
				in:   []int{8},
			},
			expect: []int{1},
		},
		"jump test position 1": test{
			// Jump test usig position mode.
			input: input{
				code: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
				in:   []int{0},
			},
			expect: []int{0},
		},
		"jump test position 2": test{
			input: input{
				code: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
				in:   []int{2},
			},
			expect: []int{1},
		},
		"jump test immediate 1": test{
			// Jump test using immediate mode.
			input: input{
				code: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
				in:   []int{0},
			},
			expect: []int{0},
		},
		"jump test immediate 2": test{
			input: input{
				code: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
				in:   []int{-1},
			},
			expect: []int{1},
		},
		"large input": test{
			input: input{
				code: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
					1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
					999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				in: []int{8},
			},
			expect: []int{1000},
		},
		"big memory": test{
			input: input{
				code: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
				in:   []int{},
			},
			expect: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		"produce big int": test{
			input: input{
				code: []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
				in:   []int{},
			},
			expect: []int{1219070632396864},
		},
		"parse big int": test{
			input: input{
				code: []int{104, 1125899906842624, 99},
				in:   []int{},
			},
			expect: []int{1125899906842624},
		},
	}

	for i, ts := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(2*time.Second))
			defer cancel()

			in := make(chan int)
			out := make(chan int)
			errs := make(chan error)

			go func(ts test) {
				c := intcode.Code(ts.input.code)
				errs <- c.Exec(ctx, in, out)
			}(ts)

			go func(ts test) {
				for _, i := range ts.input.in {
					in <- i
				}
			}(ts)

			r := []int{}
			var done bool
			for !done {
				select {
				case <-ctx.Done():
					require.NoError(t, ctx.Err())
				case i, ok := <-out:
					if !ok {
						done = true
						continue
					}
					r = append(r, i)
				case err := <-errs:
					require.NoError(t, err)
				}
			}
			require.Equal(t, ts.expect, r)
		})
	}
}
