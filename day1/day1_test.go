package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuel(t *testing.T) {
	require.Equal(t, fuel(12), 2.0)
}

func TestPart1(t *testing.T) {
	require.Equal(t, part1(), 3291356)
}
