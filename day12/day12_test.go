package main

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2019/lib"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	type test struct {
		input  string
		expect int
	}

	tests := []test{
		test{
			input: `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`,
			expect: 1940,
		},
	}

	for i, test := range tests {
		input := lib.InputInts(test.input, lib.NumberParser)
		var p, v [N]column
		for i := 0; i < N; i++ {
			p[i] = column(input[i])
			v[i] = column([]int{0, 0, 0})
		}
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			for i := 0; i < 100; i++ {
				if i%10 == 0 {
					fmt.Println("After", i, "steps")
					for x := 0; x < N; x++ {
						fmt.Println("pos", p[x], "vel", v[x])
					}
				}
				for dim := 0; dim < 3; dim++ {
					step(p, v, dim)
				}
			}
			require.Equal(t, test.expect, energy(p, v))
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
			input: `<x=-1, y=-0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>
`,
			expect: 2772,
		},
		test{
			input: `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>
`,
			expect: 4686774924,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			input := lib.InputInts(test.input, lib.NumberParser)
			require.Equal(t, test.expect, part2(input))
		})
	}
}
