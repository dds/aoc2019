package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/dds/aoc2019/lib"
	"github.com/dds/aoc2019/lib/inputs"
)

var inputRE = regexp.MustCompile(`\d+ \w+`)

var Input = lib.ParseInput(inputs.Day14(), func(s string) []string {
	return lib.TrimSpace(inputRE.FindAllString(s, -1))
})

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

func part1(input [][]string) (rc int) {
	fmt.Println(input)
	return
}

func part2(input [][]string) (rc int) {
	return
}
