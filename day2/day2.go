package main

import (
	"fmt"

	"github.com/dds/aoc2020/util"
)

var Input = util.InputInts(util.Inputs[2], util.CSVParser)[0]

func main() {
	fmt.Println(part1(Input))
	x, y := part2(Input)
	fmt.Println(100*x + y)
}

type Opcode int

const (
	Unknown Opcode = iota
	Add
	Mul
	Halt = 99
)

func Exec(code []int) ([]int, error) {
	i := 0
	for op := Opcode(code[i]); op != Halt; op = Opcode(code[i]) {
		args := [3]int{code[i+1], code[i+2], code[i+3]}
		switch op {
		default:
			return nil, fmt.Errorf("Unknown op: %v, i: %v, code: %v", op, i, code)
		case Add:
			code[args[2]] = code[args[0]] + code[args[1]]
			i += 4
		case Mul:
			code[args[2]] = code[args[0]] * code[args[1]]
			i += 4
		}
	}
	return code, nil
}

func part1(i []int) int {
	code, _ := Try(i, 12, 2)
	return code[0]
}

func Try(input []int, noun, verb int) ([]int, error) {
	i := make([]int, len(input))
	copy(i, input)
	i[1] = noun
	i[2] = verb
	r, err := Exec(i)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func part2(input []int) (x int, y int) {
	result := 19690720
	for x = 0; x < 100; x++ {
		for y = 0; y < 100; y++ {
			r, err := Try(input, x, y)
			if err != nil {
				panic(fmt.Errorf("%w: x: %v, y: %v", err, x, y))
			}
			if r[0] == result {
				return
			}
		}
	}
	return
}
