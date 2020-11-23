package util_test

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2020/util"
	"github.com/stretchr/testify/require"
)

func TestGrid(t *testing.T) {
	g := &util.Grid{}
	require.Equal(t, ".x\n..\n", fmt.Sprint(g.AddPoint(1, 1, "x")))
	require.Equal(t, "..\nxx\n", fmt.Sprint(g.AddStrip(0, 0, 2, 'R', "x")))
}
