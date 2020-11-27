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
			input: `10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL
`,
			expect: 31,
		},
		test{
			input: `9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL`,
			expect: 165,
		},
		test{
			input: `157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`,
			expect: 13312,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			in := lib.ParseInput(test.input, Parser)
			require.Equal(t, test.expect, part1(in))
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
			input: `157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`,
			expect: 82892753,
		},
		// test{
		// 	input: ``,
		// 	expect: ,
		// },
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			in := lib.ParseInput(test.input, Parser)
			require.Equal(t, test.expect, part2(in))
		})
	}
}
