package main

import (
	"fmt"
	"regexp"

	"github.com/dds/aoc2019/lib"
	"github.com/dds/aoc2019/lib/inputs"
)

func parse(input string) []int {
	return lib.InputInts(input, func(input string) []string {
		return regexp.MustCompile(`\d`).FindAllString(input, -1)
	})[0]
}

var Input = parse(inputs.Day16())

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
