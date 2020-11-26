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
			input: `9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL
`,
			expect: 165,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			in := lib.ParseInput(test.input, Parser)
			m := mkformulae(in)
			require.Equal(t, test.expect, m.ore("FUEL"))
		})
	}
}
