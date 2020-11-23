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

	tests := []test{
		test{
			input: input{
				code: []int{3, 0, 4, 0, 99},
				in:   []int{1},
			},
			expect: result{
				code: []int{1, 0, 4, 0, 99},
				out:  []int{1},
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

	for _, test := range tests {
		r, _, err := intcode.Exec(test.input, []int{})
		require.NoError(t, err)
		require.Equal(t, test.expect, r)
	}
}
