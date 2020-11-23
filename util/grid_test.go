package util_test

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2020/util"
	"github.com/stretchr/testify/require"
)

func TestGrid(t *testing.T) {
	g := &util.Grid{}
	g = g.AddPoint(1, 1, "X")
	require.Equal(t, fmt.Sprint(g), ".X\n..\n")
}
