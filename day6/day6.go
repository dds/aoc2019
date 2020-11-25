package main

import (
	"fmt"
	"strings"

	"github.com/dds/aoc2019/lib"
)

const (
	COM = "COM"
	YOU = "YOU"
	SAN = "SAN"
)

var Input = lib.ParseInput(lib.Inputs[6], ParseInput)

func ParseInput(input string) []string {
	return lib.TrimSpace(strings.FieldsFunc(input, func(c rune) bool { return c == ')' }))
}

func main() {
	fmt.Println(Part1(Input))
	fmt.Println(Part2(Input))
}

func Part1(input [][]string) (r int) {
	return CountOrbits(Orbits(input))
}

func CountOrbits(o map[string]string) (r int) {
	for _, v := range o {
		next := v
		for next != COM {
			r++
			next = o[next]
		}
	}
	r += len(o)
	return
}

func Orbits(input [][]string) map[string]string {
	r := map[string]string{}
	for _, fields := range input {
		r[fields[1]] = fields[0]
	}
	return r
}

func Part2(input [][]string) (r int) {
	return CountTransfers(Orbits(input))
}

func CountTransfers(o map[string]string) int {
	var (
		YourPath, SantasPath []string
		next                 = o[YOU]
	)

	for next != "" {
		YourPath = append(YourPath, o[next])
		next = o[next]
	}

	next = o[SAN]
	for next != "" {
		SantasPath = append(SantasPath, o[next])
		next = o[next]
	}

	for i := 0; i < len(YourPath); i++ {
		for j := 0; j < len(SantasPath); j++ {
			if YourPath[i] == SantasPath[j] {
				return i + j + 2
			}
		}
	}
	return 0
}
