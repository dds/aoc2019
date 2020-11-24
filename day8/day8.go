package main

import (
	"fmt"
	"strings"

	"github.com/dds/aoc2020/util"
)

var Input = strings.TrimSpace(util.Inputs[8])

func main() {
	fmt.Println(part1(Input))
}

const (
	w = 25
	h = 6
)

func part1(input string) (r string) {
	fmt.Println(len(Input), "input size")
	fmt.Println(w*h, "W * H")
	fmt.Println(len(Input)/(w*h), "input / size = frames")
	m := map[int]string{}
	var minZeroIdx int
	minZero := 1 << 32
	for i := 0; i < len(input); i += w * h {
		zeros := 0
		for j := 0; j < w*h; j++ {
			if string(input[i+j]) == "0" {
				zeros++
			}
		}
		if zeros < minZero {
			minZero = zeros
			minZeroIdx = i
		}
		m[i] = input[i : i+w*h]
	}
	fmt.Println(minZeroIdx, "minZeroIdx")

	var ones, twos int
	for i := 0; i < 150; i++ {
		switch s := string(m[minZeroIdx][i]); s {
		case "1":
			ones++
		case "2":
			twos++
		}
	}
	r = fmt.Sprint(ones * twos)
	return
}
