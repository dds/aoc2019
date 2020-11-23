package intcode_test

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2020/intcode"
	"github.com/dds/aoc2020/util"
	"github.com/stretchr/testify/require"
)

func TestIO(t *testing.T) {
	type input struct {
		code, in []int
	}
	type result struct {
		code, out []int
	}
	type test struct {
		input  input
		expect result
	}

	tests := map[string]test{
		"1st": test{
			input: input{
				code: []int{3, 0, 4, 0, 99},
				in:   []int{1},
			},
			expect: result{
				code: []int{1, 0, 4, 0, 99},
				out:  []int{1},
			},
		},
		"2nd": test{
			input: input{
				code: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
				in:   []int{8},
			},
			expect: result{
				code: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8},
				out:  []int{1},
			},
		},
		"3rd": test{
			input: input{
				code: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
				in:   []int{9},
			},
			expect: result{
				code: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8},
				out:  []int{0},
			},
		},
		"4th": test{
			input: input{
				code: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
				in:   []int{7},
			},
			expect: result{
				code: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8},
				out:  []int{1},
			},
		},
		"5th": test{
			input: input{
				code: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
				in:   []int{8},
			},
			expect: result{
				code: []int{3, 3, 1108, 1, 8, 3, 4, 3, 99},
				out:  []int{1},
			},
		},
		"jump test position 1": test{
			// Jump test usig position mode.
			input: input{
				code: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
				in:   []int{0},
			},
			expect: result{
				code: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 0, 0, 1, 9},
				out:  []int{0},
			},
		},
		"jump test position 2": test{
			input: input{
				code: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
				in:   []int{2},
			},
			expect: result{
				code: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 2, 1, 1, 9},
				out:  []int{1},
			},
		},
		"jump test immediate 1": test{
			// Jump test using immediate mode.
			input: input{
				code: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
				in:   []int{0},
			},
			expect: result{
				code: []int{3, 3, 1105, 0, 9, 1101, 0, 0, 12, 4, 12, 99, 0},
				out:  []int{0},
			},
		},
		"jump test immediate 2": test{
			input: input{
				code: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
				in:   []int{-1},
			},
			expect: result{
				code: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
				out:  []int{1},
			},
		},
		"large input": test{
			input: input{
				code: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
					1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
					999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				in: []int{8},
			},
			expect: result{
				code: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 1000, 8, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				out:  []int{1000},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c, out, err := intcode.Exec(test.input.code, test.input.in)
			require.NoError(t, err)
			require.Equal(t, test.expect.code, c)
			require.Equal(t, test.expect.out, out)
		})
	}
}

func TestOpmodes(t *testing.T) {
	type test struct {
		input  int
		expect []intcode.Opmode
	}

	tests := []test{
		test{
			input:  1002,
			expect: []intcode.Opmode{intcode.PositionMode, intcode.ImmediateMode, intcode.PositionMode},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, intcode.Opmodes(test.input))
		})
	}

}

func TestOps(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}
	tests := []test{
		test{
			input: util.InputInts("1,9,10,3,2,3,11,0,99,30,40,50", util.CSVParser)[0],
			expect: []int{
				3500, 9, 10, 70,
				2, 3, 11, 0,
				99,
				30, 40, 50,
			},
		},
		test{
			input: util.InputInts("1,0,0,0,99", util.CSVParser)[0],
			expect: []int{
				2, 0, 0, 0, 99,
			},
		},
		test{
			input: util.InputInts("2,3,0,3,99", util.CSVParser)[0],
			expect: []int{
				2, 3, 0, 6, 99,
			},
		},
		test{
			input: util.InputInts("2,4,4,5,99,0", util.CSVParser)[0],
			expect: []int{
				2, 4, 4, 5, 99, 9801,
			},
		},
		test{
			input: util.InputInts("1,1,1,4,99,5,6,0,99", util.CSVParser)[0],
			expect: []int{
				30, 1, 1, 4, 2, 5, 6, 0, 99,
			},
		},
		test{
			input: util.InputInts("1002,4,3,4,33", util.CSVParser)[0],
			expect: []int{
				1002, 4, 3, 4, 99,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			r, _, err := intcode.Exec(test.input, []int{})
			require.NoError(t, err)
			require.Equal(t, test.expect, r)
		})
	}
}
