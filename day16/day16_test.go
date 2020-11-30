package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	type test struct {
		input  string
		expect int
	}

	tests := []test{
		test{
			input:  "12345678",
			expect: 23845678,
		},
		test{
			input:  "80871224585914546619083218645595",
			expect: 24176176,
		},
		test{
			input:  "19617804207202209144916044189917",
			expect: 73745418,
		},
		test{
			input:  "69317163492948606335995924319873",
			expect: 52432133,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, part1(parse(test.input)))
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
			input:  "12345678",
			expect: 23845678,
		},
		test{
			input:  "80871224585914546619083218645595",
			expect: 24176176,
		},
		test{
			input:  "19617804207202209144916044189917",
			expect: 73745418,
		},
		test{
			input:  "69317163492948606335995924319873",
			expect: 52432133,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, test.expect, part2(parse(test.input)))
		})
	}
}
