package main

import (
	"testing"

	"github.com/dds/aoc2019/lib"
	"github.com/stretchr/testify/require"
)

func TestOps(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}
	tests := []test{
		test{
			input: lib.InputInts("1,9,10,3,2,3,11,0,99,30,40,50", lib.CSVParser)[0],
			expect: []int{
				3500, 9, 10, 70,
				2, 3, 11, 0,
				99,
				30, 40, 50,
			},
		},
		test{
			input: lib.InputInts("1,0,0,0,99", lib.CSVParser)[0],
			expect: []int{
				2, 0, 0, 0, 99,
			},
		},
		test{
			input: lib.InputInts("2,3,0,3,99", lib.CSVParser)[0],
			expect: []int{
				2, 3, 0, 6, 99,
			},
		},
		test{
			input: lib.InputInts("2,4,4,5,99,0", lib.CSVParser)[0],
			expect: []int{
				2, 4, 4, 5, 99, 9801,
			},
		},
		test{
			input: lib.InputInts("1,1,1,4,99,5,6,0,99", lib.CSVParser)[0],
			expect: []int{
				30, 1, 1, 4, 2, 5, 6, 0, 99,
			},
		},
	}

	for _, test := range tests {
		r, err := Exec(test.input)
		require.NoError(t, err)
		require.Equal(t, test.expect, r)
	}
}
