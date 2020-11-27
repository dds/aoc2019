package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/dds/aoc2019/lib"
	"github.com/dds/aoc2019/lib/inputs"
)

var inputRE = regexp.MustCompile(`\d+ \w+`)

var Input = lib.ParseInput(inputs.Day14(), Parser)

func Parser(s string) []string {
	return lib.TrimSpace(inputRE.FindAllString(s, -1))
}

func main() {
	fmt.Println(part1(Input))
	fmt.Println(part2(Input))
}

const (
	Ore  = "ORE"
	Fuel = "FUEL"
)

type reagent struct {
	n   int
	typ string
}

func read(s string) (t reagent) {
	// From "5 NZVS" into reagent{n: 5, typ: "NZVS"}
	i := strings.Fields(s)
	n, err := strconv.Atoi(i[0])
	if err != nil {
		panic(err)
	}
	t.n = n
	t.typ = i[1]
	return
}

// A formula produces a number of output from a list of reagents.
type formula struct {
	outputs  int
	reagents []reagent
}
type formulae map[string]formula

func mkformulae(input [][]string) formulae {
	m := make(formulae)
	for _, row := range input {
		t := read(row[len(row)-1])
		f := formula{outputs: t.n}
		for j := 0; j < len(row)-1; j++ {
			f.reagents = append(f.reagents, read(row[j]))
		}
		m[t.typ] = f
	}
	return m
}

func (f formulae) ore(typ string, n int, byproducts map[string]int) int {
	var ore = 0
	var required = []reagent{reagent{n: n, typ: typ}}
	for len(required) > 0 {
		t := required[0]
		required = required[1:]
		var n = t.n
		if t.typ == Ore {
			ore += n
			continue
		}
		if byproducts[t.typ] > n {
			byproducts[t.typ] -= n
		} else {
			n -= byproducts[t.typ]
			byproducts[t.typ] = 0
		}
		recipe := f[t.typ]
		scale := int(math.Ceil(float64(n) / float64(recipe.outputs)))
		byproducts[t.typ] += scale*recipe.outputs - n
		for _, reagent := range recipe.reagents {
			reagent.n = scale * reagent.n
			required = append(required, reagent)
		}
	}
	return ore
}

func part1(input [][]string) (rc int) {
	m := mkformulae(input)
	rc = m.ore(Fuel, 1, map[string]int{})
	return
}

func part2(input [][]string) (rc int) {
	m := mkformulae(input)
	ore := 1000 * 1000 * 1000 * 1000
	var i = 1
	rc = 1
	for {
		o := m.ore(Fuel, rc+i, map[string]int{})
		if o > ore {
			if i != 1 {
				i = 1
				continue
			}
			break
		}
		rc += i
		i *= 2
	}
	return
}
