package intcode_test

import (
	"fmt"
	"testing"

	"github.com/dds/aoc2020/intcode"
	"github.com/stretchr/testify/require"
)

func TestIO(t *testing.T) {
	type input struct {
		code, in []int
	}
	type result struct {
		code, out []int
	}
	type test struct {
		input  input
		expect result
	}

	tests := []test{
		test{
			input: input{
				code: []int{3, 0, 4, 0, 99},
				in:   []int{1},
			},
			expect: result{
				code: []int{1, 0, 4, 0, 99},
				out:  []int{1},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c, out, err := intcode.Exec(test.input.code, test.input.in)
			require.NoError(t, err)
			require.Equal(t, test.expect.code, c)
			require.Equal(t, test.expect.out, out)
		})
	}
}
