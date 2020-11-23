package main

import (
	"fmt"
	"strconv"

	"github.com/dds/aoc2020/util"
)

var Input = util.ParseInput(util.Inputs[3], util.CSVParser)[0]

func Parse(pos string) (rune, int) {
	dir := pos[0]
	steps, err := strconv.Atoi(pos[1:])
	if err != nil {
		panic(err)
	}
	return rune(dir), steps
}

func Cross(path1 []string, path2 []string) int {
	g := &util.Grid{}
	var x, y int
	for _, i := range path1 {
		dir, steps := Parse(i)
		x, y = g.Translate(g.Walk(x, y, steps, dir))
		g = g.AddStrip(x, y, steps, dir, "-")
	}
	for _, i := range path2 {
		dir, steps := Parse(i)
		x, y = g.Translate(g.Walk(x, y, steps, dir))
		g = g.AddStrip(x, y, steps, dir, "-")
	}
	fmt.Println(g)
	return 0
}

func main() {

}
