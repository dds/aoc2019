package main

import (
	"fmt"
	"math"

	"github.com/dds/aoc2020/util"
)

func main() {
	input := util.InputNums(1, util.CSVParser)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input [][]float64) int {
	sum := 0.0
	for _, l := range input {
		sum += fuel(l[0])
	}
	return int(sum)
}

func fuel(mass float64) float64 {
	return math.Floor(mass/3.0) - 2
}

func part2(input [][]float64) int {
	mass := 0.0
	for _, l := range input {
		mass += massfuel(l[0])
	}
	return int(mass)
}

func massfuel(mass float64) float64 {
	sum := 0.0
	for f := fuel(mass); f >= 0; f = fuel(f) {
		sum += f
	}
	return sum
}
