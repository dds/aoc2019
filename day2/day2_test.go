package main

import (
	"testing"

	"github.com/dds/aoc2020/util"
	"github.com/stretchr/testify/require"
)

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
	}

	for _, test := range tests {
		require.Equal(t, test.expect, Exec(test.input))
	}
}
