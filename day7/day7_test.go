package main

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/dds/aoc2020/intcode"
	"github.com/stretchr/testify/require"
)

func TestParsePhase(t *testing.T) {
	type result struct {
		out []int
		err error
	}
	type test struct {
		phase  int
		expect result
	}

	tests := []test{
		test{
			phase: 12034,
			expect: result{
				out: []int{1, 2, 0, 3, 4},
			},
		},
		test{
			phase: 2134,
			expect: result{
				out: []int{0, 2, 1, 3, 4},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			a, err := parsePhase(test.phase)
			if err != nil {
				require.True(t, errors.Is(err, test.expect.err), err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expect.out, a)
		})
	}

}
func TestPart1(t *testing.T) {
	type input struct {
		code, in []int
	}
	type test struct {
		input  input
		expect int
	}

	tests := []test{
		test{
			input: input{
				code: []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
				in:   []int{4, 3, 2, 1, 0},
			},
			expect: 43210,
		},
		test{
			input: input{
				code: []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
					101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
				in: []int{0, 1, 2, 3, 4},
			},
			expect: 54321,
		},
		test{
			input: input{
				code: []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
					1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
				in: []int{1, 0, 4, 3, 2},
			},
			expect: 65210,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(2*time.Second))
			defer cancel()

			code := intcode.Code(test.input.code)

			require.Equal(t, test.expect, Part1(ctx, code, test.input.in))
		})
	}
}

func TestPart2(t *testing.T) {
	type input struct {
		code, in []int
	}
	type test struct {
		input  input
		expect int
	}

	tests := []test{
		test{
			input: input{
				code: []int{3, 26,
					1001, 26, -4, 26,
					3, 27,
					1002, 27, 2, 27,
					1, 27, 26, 27,
					4, 27,
					1001, 28, -1, 28,
					1005, 28, 6,
					99,
					0, 0, 5},
				in: []int{9, 8, 7, 6, 5},
			},
			expect: 139629729,
		},
		test{
			input: input{
				code: []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54,
					-5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4,
					53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
				in: []int{9, 7, 8, 5, 6},
			},
			expect: 18216,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			ctx, cancel := context.WithDeadline(context.TODO(), time.Now().Add(2*time.Second))
			defer cancel()

			code := intcode.Code(test.input.code)

			require.Equal(t, test.expect, Part2(ctx, code, test.input.in))
		})
	}
}
