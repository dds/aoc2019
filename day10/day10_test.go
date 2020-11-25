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
		test{
			input: `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..
`,
			score: 41,
			point: point{6, 3},
			// ...
		},
		test{
			input: `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
			score: 202,
			point: point{11, 13},
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

func TestPart2(t *testing.T) {
	type test struct {
		input        string
		base, target point
	}

	tests := []test{
		test{
			input: `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`,
			base:   point{11, 13},
			target: point{11, 12},
			// ...
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			field := Field(Parse(test.input))
			targets := vaporize(test.base, field)
			err := fmt.Sprintf("\n\ttargets (%d): %v\n\n\tfield (%d): %v", len(targets), targets, len(field), field)
			require.Equal(t, test.target, targets[0].point, err)
		})
	}
}
