package main

import (
	"fmt"
	"strconv"

	"github.com/dds/aoc2019/lib"
)

var Input = lib.InputInts(lib.Inputs[4], lib.DashParser)[0]

func main() {
	min := Input[0]
	max := Input[1]
	fmt.Println(part1(min, max))
	fmt.Println(part2(min, max))
}

func part1(min, max int) int {
	sum := 0
	for i := min; i < max; i++ {
		if password1(i) {
			sum += 1
		}
	}
	return sum
}

func password1(input int) bool {
	s := fmt.Sprint(input)
	var adj bool
	var dec bool
	for i := 0; i+1 < len(s); i++ {
		if s[i] == s[i+1] {
			adj = true
		}
		t, err := strconv.Atoi(string(s[i]))
		if err != nil {
			panic(err)
		}
		v, err := strconv.Atoi(string(s[i+1]))
		if err != nil {
			panic(err)
		}
		if t > v {
			dec = true
		}
	}
	if !adj {
		return false
	}
	if dec {
		return false
	}
	return true
}

func part2(min, max int) int {
	sum := 0
	for i := min; i < max; i++ {
		if password2(i) {
			sum += 1
		}
	}
	return sum
}

func password2(input int) bool {
	s := fmt.Sprint(input)
	var dec bool
	for i := 0; i < 5; i++ {
		t, err := strconv.Atoi(string(s[i]))
		if err != nil {
			panic(err)
		}
		v, err := strconv.Atoi(string(s[i+1]))
		if err != nil {
			panic(err)
		}
		if t > v {
			dec = true
		}
	}
	if dec {
		return false
	}
	for i := 0; i < 5; i++ {
		if i < 4 && s[i] == s[i+1] && s[i] == s[i+2] {
			continue
		} else if i > 0 && s[i-1] == s[i] && s[i] == s[i+1] {
			continue
		} else if s[i] == s[i+1] {
			return true
		}
	}
	return false
}
