package main

import (
	"fmt"

	"github.com/dds/aoc2020/intcode"
	"github.com/dds/aoc2020/util"
)

var Input = util.InputInts(util.Inputs[5], util.CSVParser)[0]

func main() {
	fmt.Println(part1(Input))
}

func part1(input []int) (r []int) {
	var err error
	_, r, err = intcode.Exec(input, []int{1})
	if err != nil {
		panic(err)
	}
	return
}
