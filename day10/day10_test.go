package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	type test struct {
		input string
		score int
		point point
	}

	tests := []test{
		test{
			input: `.#..#
.....
#####
....#
...##
`,
			score: 8,
			point: point{3, 4},
			// ...
		},
		test{
			input: `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####
`,
			score: 33,
			point: point{5, 8},
			// ...
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			score, p := part1(test.input)
			require.Equal(t, test.point, p)
			require.Equal(t, test.score, score)
		})
	}
}
