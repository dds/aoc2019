package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFuel(t *testing.T) {
	require.Equal(t, fuel(12), 2.0)
}
