package main

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2019/util"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	type test struct {
		input  string
		expect int
	}
	tests := []test{
		test{
			input: `R8,U5,L5,D3
U7,R6,D4,L4`,
			expect: 6,
		},
		test{
			input: `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			expect: 159,
		},
		test{
			input: `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			expect: 135,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			input := util.ParseInput(test.input, util.CSVParser)
			path1 := input[0]
			path2 := input[1]
			require.Equal(t, test.expect, Cross(path1, path2))
		})
	}
}

func TestPart2(t *testing.T) {
	type test struct {
		input  string
		expect int
	}
	tests := []test{
		test{
			input: `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`,
			expect: 610,
		},
		test{
			input: `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`,
			expect: 410,
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			input := util.ParseInput(test.input, util.CSVParser)
			path1 := input[0]
			path2 := input[1]
			require.Equal(t, test.expect, MinWalk(path1, path2))
		})
	}
}
