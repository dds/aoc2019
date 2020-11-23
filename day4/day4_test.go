package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	type test struct {
		input  int
		expect bool
	}

	tests := []test{
		test{
			input:  111111,
			expect: true,
		},
		test{
			input:  223450,
			expect: false,
		},
		test{
			input:  123789,
			expect: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, password1(test.input))
		})
	}
}

func TestPart2(t *testing.T) {
	type test struct {
		input  int
		expect bool
	}

	tests := []test{
		test{
			input:  112233,
			expect: true,
		},
		test{
			input:  123444,
			expect: false,
		},
		test{
			input:  111122,
			expect: true,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, password2(test.input))
		})
	}
}
