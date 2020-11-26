package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

const Ore = "ORE"

type rec struct {
	n   int
	typ string
}

func read(s string) (r rec) {
	i := strings.Fields(s)
	n, err := strconv.Atoi(i[0])
	if err != nil {
		panic(err)
	}
	r.n = n
	r.typ = i[1]
	return
}

func part1(input [][]string) (rc int) {
	m := map[rec][]rec{}
	for _, row := range input {
		last := read(row[len(row)-1])
		m[last] = []rec{}
		for j := 0; j < len(row)-1; j++ {
			m[last] = append(m[last], read(row[j]))
		}
	}
	fmt.Println(m)
	return
}

func part2(input [][]string) (rc int) {
	return
}
