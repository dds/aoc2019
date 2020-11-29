package main

import (
	"fmt"
	"regexp"
	"strconv"

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

var base = []int{0, 1, 0, -1}

func coefficients(n, row int) []int {
	r := make([]int, n+1)
	var i, digit int
	for i < n+1 {
		for j := 0; j < row+1 && i+j < n+1; j++ {
			r[i+j] = base[digit]
		}
		i += row + 1
		digit = (digit + 1) % 4
	}
	r = r[1:]
	return r
}

func phase(input []int) (output []int) {
	n := len(input)
	output = make([]int, n)
	for i := 0; i < n; i++ {
		coef := coefficients(n, i)
		var c int
		for j := 0; j < len(coef); j++ {
			c += input[j] * coef[j]
		}
		if c < 0 {
			c = -c
		}
		output[i] = c % 10
	}
	return
}

func part1(input []int) (rc int) {
	fmt.Println("Input len: ", len(input))
	fmt.Println("Initial input: ", input)
	fmt.Println("first coefficients: ", coefficients(len(input), 0))
	fmt.Println("second coefficients: ", coefficients(len(input), 1))
	phased := phase(input)
	fmt.Println("First phase:", phased)

	for i := 0; i < 99; i++ {
		phased = phase(phased)
	}
	var join string
	for _, i := range phased {
		join += fmt.Sprint(i)
	}
	a, err := strconv.Atoi(join[:8])
	if err != nil {
		panic(err)
	}
	return a
}

func part2(input []int) (rc int) {
	signal := phase(input)
	for i := 0; i < 999; i++ {
		signal = phase(signal)
	}

	var join string
	for _, i := range signal {
		join += fmt.Sprint(i)
	}
	a, err := strconv.Atoi(join[:8])
	if err != nil {
		panic(err)
	}
	return a
}
