package main

import (
	"fmt"
	"math"

	"github.com/dds/aoc2020/util"
)

func fuel(mass float64) float64 {
	return math.Floor(mass/3.0) - 2
}

func part1() int {
	input, err := util.InputNums(1, util.CSVParser)
	if err != nil {
		panic(err)
	}
	sum := 0.0
	for _, l := range input {
		sum += fuel(l[0])
	}
	return int(sum)
}

func massfuel(mass float64) float64 {
	sum := 0.0
	for f := fuel(mass); f >= 0; f = fuel(f) {
		sum += f
	}
	return sum
}

func part2() int {
	input, err := util.InputNums(1, util.CSVParser)
	if err != nil {
		panic(err)
	}
	mass := 0.0
	for _, l := range input {
		mass += massfuel(l[0])
	}
	return int(mass)
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
