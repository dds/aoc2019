package main

import (
	"fmt"
	"strconv"

	"github.com/dds/aoc2020/util"
)

var Input = util.ParseInput(util.Inputs[3], util.CSVParser)[0]

func ParsePos(pos string) (int, int) {
	dir := pos[0]
	steps, err := strconv.Atoi(pos[1:])
	if err != nil {
		panic(err)
	}
	switch dir {
	case 'U':
		return 0, steps
	case 'D':
		return 0, -steps
	case 'L':
		return -steps, 0
	case 'R':
		return steps, 0
	}
	return 0, 0
}

func Cross(path1 []string, path2 []string) int {
	g := &util.Grid{}
	for _, i := range path1 {
		x, y := ParsePos(i)
		g = g.AddPoint(x, y, "X")
	}
	fmt.Println(g)
	return 0
}

func main() {

}
