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

type coefficients []int

func (c coefficients) String() (s string) {
	s = "["
	for _, i := range c {
		switch i {
		case 0:
			s += "0"
		case 1:
			s += "1"
		case -1:
			s += "!"
		}
		s += " "
	}
	s = s[:len(s)-1] + "]"
	return
}

func mkcoefficients(n, row int) coefficients {
	r := make(coefficients, n+1)
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

func phase(input []int, offset int) (output []int) {
	n := len(input)
	output = make([]int, n)
	for i := 0; i < n; i++ {
		coef := mkcoefficients(n, i)
		var c int
		for j := 0; j < len(coef); j++ {
			c += input[offset+j] * coef[offset+j]
		}
		if c < 0 {
			c = -c
		}
		output[i] = c % 10
	}
	return
}

func phase2ndhalfonly(input []int, offset int) (output []int) {
	n := len(input)
	output = make([]int, n)
	output[n-1] = input[n-1]
	for i := n - 2; i > -1; i-- {
		output[i] = output[i+1] + input[i]
		output[i] %= 10
	}
	return
}

func part1(input []int) (rc int) {
	n := len(input)
	m := lib.Min(n, 32)
	fmt.Println("N: ", n)
	fmt.Println("    input:", input[:m])
	for i := 0; i < lib.Min(n, 32); i++ {
		fmt.Printf("%02vth mult: %v\n", i+1, mkcoefficients(n, i)[:m])
	}
	phased := phase(input, 0)
	for i := 0; i < 99; i++ {
		fmt.Printf("%02vth phse: %v\n", i+1, phased[:m])
		phased = phase(phased, 0)
	}

	fmt.Println("Nth phase:", phased[:m])
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

func part2(signal []int) (rc int) {
	n := len(signal)
	input := make([]int, 10000*n)
	for i := 0; i < 10000; i++ {
		copy(input[n*i:n*(i+1)], signal)
	}

	var s string
	for i := 0; i < 7; i++ {
		s += fmt.Sprint(signal[i])
	}
	offset, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	phased := phase2ndhalfonly(input, 0)
	for i := 0; i < 99; i++ {
		phased = phase2ndhalfonly(phased, 0)
	}

	s = ""
	for i := 0; i < 8; i++ {
		s += fmt.Sprint(phased[offset+i])
	}
	rc, err = strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return
}
