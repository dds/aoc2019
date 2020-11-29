package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/dds/aoc2019/lib"
	"github.com/dds/aoc2019/lib/inputs"
)

var Input = lib.InputInts(inputs.Day16(), func(input string) []string {
	return regexp.MustCompile(`\d`).FindAllString(input, -1)
})[0]

func Test(t *testing.T) {
	// type test struct {
	// 	input  int
	// 	expect int
	// }

	// tests := []test{
	// 	test{
	// 		// ...
	// 	},
	// }

	// for i, test := range tests {
	// 	t.Run(fmt.Sprint(i), func(t *testing.T) {
	// 		require.Equal(t, test.expect, test.input)
	// 	})
	// }
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}

func part1(input []int) (rc int) {
	fmt.Println(input)
	return
}

func part2(input []int) (rc int) {
	return
}
