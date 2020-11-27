package main

import (
	"fmt"
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

type term struct {
	n   int
	typ string
}

func read(s string) (t term) {
	// From "5 NZVS" into term{n: 5, typ: "NZVS"}
	i := strings.Fields(s)
	n, err := strconv.Atoi(i[0])
	if err != nil {
		panic(err)
	}
	t.n = n
	t.typ = i[1]
	return
}

// A formula produces a number of output from a list of terms.
type formula struct {
	outputs int
	terms   []term
	ore     float64
}
type formulae map[string]formula

func mkformulae(input [][]string) formulae {
	m := make(formulae)
	for _, row := range input {
		t := read(row[len(row)-1])
		f := formula{outputs: t.n}
		for j := 0; j < len(row)-1; j++ {
			f.terms = append(f.terms, read(row[j]))
		}
		m[t.typ] = f
	}
	// Add ore amounts for fundamental types.
	for _, f := range m {
		if len(f.terms) == 1 && f.terms[0].typ == Ore {
			f.ore = float64(f.terms[0].n) / float64(f.outputs)
		}
	}
	return m
}

func (f formulae) fuelOre(n int) int {
	m := map[string]int{} // By products
	formula := f[Fuel]
	terms := make([]term, len(formula.terms))
	copy(terms, formula.terms)
	ore := 0

	for _, t := range terms {
		t.n *= n
		var o int
		o, m = f.ore(t, m)
		ore += o
	}
	return ore
}

func (f formulae) ore(t term, m map[string]int) (int, map[string]int) {

	ore := 0

	for m[t.typ] < t.n {
		for _, trm := range f[t.typ].terms {
			if trm.typ == Ore {
				ore += trm.n
				continue
			}
			if m[trm.typ] > trm.n {
				m[trm.typ] -= trm.n
				continue
			}
			var o int
			o, m = f.ore(trm, m)
			ore += o
			m[trm.typ] -= trm.n
		}
		m[t.typ] += f[t.typ].outputs
	}
	return ore, m
}

func part1(input [][]string) (rc int) {
	m := mkformulae(input)
	rc = m.fuelOre(1)
	return
}

func part2(input [][]string) (rc int) {
	m := mkformulae(input)
	orePerFuel := m.fuelOre(1)
	maxOre := 1000 * 1000 * 1000 * 1000
	oneThousand := m.fuelOre(1000)
	var a = maxOre / orePerFuel
	fmt.Println("One FUEL is ", orePerFuel, "ore. Estimating ", a, "min fuel.")
	fmt.Println("1000 FUEL is ", oneThousand, "ore.")

	var ore = 0
	for i := 0; i < 2440; i++ {
		ore += oneThousand
	}
	c := 2440 * 1000
	fmt.Println(c, "FUEL is ", ore, "ore.")
	var j = 0
	for ore < maxOre {
		ore += m.fuelOre(1)
		j++
	}
	fmt.Println(j, "j, ", c+j, "FUEL is ", ore, "ore.")
	rc = c + j
	return
}
