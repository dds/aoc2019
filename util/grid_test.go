package util_test

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2020/util"
	"github.com/stretchr/testify/require"
)

func TestGrid(t *testing.T) {
	type test struct {
		input  func() *util.Grid
		expect string
	}
	tests := []test{
		test{
			input: func() *util.Grid {
				g := &util.Grid{}
				g.AddPoint(1, 1, "x")
				return g
			},
			expect: ".x\n..\n",
		},
		test{
			input: func() *util.Grid {
				g := &util.Grid{}
				g.AddStrip(0, 0, 2, 'R', "x")
				return g
			},
			expect: "..\nxx\n",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, fmt.Sprint(test.input()))
		})
	}
}
