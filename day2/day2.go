package main

import (
	"fmt"

	"github.com/dds/aoc2020/util"
)

func main() {
	input := util.InputInts(util.Inputs[2], util.CSVParser)

	fmt.Println(part1(input))
}

type opcode int

const (
	Unknown opcode = iota
	Add
	Mul
	Halt = 99
)

func part1(i [][]int) int {
	return i[0][0]
}
