package main

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2020/util"
	"github.com/stretchr/testify/require"
)

func TestOrbits(t *testing.T) {
	type result struct {
		m map[string]string
		n int
		t int
	}
	type test struct {
		input  [][]string
		expect result
	}

	tests := []test{
		test{
			input: util.ParseInput(`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
`, ParseInput),
			expect: result{
				n: 42,
				m: map[string]string{
					"B": "COM",
					"G": "B",
					"H": "G",
					"C": "B",
					"D": "C",
					"I": "D",
					"E": "D",
					"J": "E",
					"K": "J",
					"L": "K",
					"F": "E",
				},
			},
		},
		test{
			input: util.ParseInput(`COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`, ParseInput),
			expect: result{
				n: 54,
				m: map[string]string{
					"B":   "COM",
					"G":   "B",
					"H":   "G",
					"C":   "B",
					"D":   "C",
					"I":   "D",
					"SAN": "I",
					"E":   "D",
					"J":   "E",
					"K":   "J",
					"YOU": "K",
					"L":   "K",
					"F":   "E",
				},
				t: 4,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			o := Orbits(test.input)
			require.Equal(t, test.expect.m, o)
			require.Equal(t, test.expect.n, CountOrbits(o))
			if o["YOU"] != "" && o["SAN"] != "" {
				require.Equal(t, test.expect.t, CountTransfers(o))
			}
		})
	}
}
