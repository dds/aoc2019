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

func main() {
	fmt.Println(part1())
}
