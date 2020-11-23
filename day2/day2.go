package main

import (
	"fmt"

	"github.com/dds/aoc2020/util"
)

var Input = util.InputInts(util.Inputs[2], util.CSVParser)[0]

func main() {
	fmt.Println(part1(Input))
}

type Opcode int

const (
	Unknown Opcode = iota
	Add
	Mul
	Halt = 99
)

func Exec(code []int) []int {
	i := 0
	for op := Opcode(code[i]); op != Halt; op = Opcode(code[i]) {
		args := [3]int{code[i+1], code[i+2], code[i+3]}
		switch op {
		default:
			panic(fmt.Errorf("Unknown op: %v", op))
		case Add:
			code[args[2]] = code[args[0]] + code[args[1]]
		case Mul:
			code[args[2]] = code[args[0]] * code[args[1]]
		}
		i += 4
	}
	return code
}

func part1(i []int) int {
	i[1] = 12
	i[2] = 2
	return Exec(i)[0]
}
