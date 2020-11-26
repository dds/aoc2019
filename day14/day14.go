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
	i := strings.Fields(s)
	n, err := strconv.Atoi(i[0])
	if err != nil {
		panic(err)
	}
	t.n = n
	t.typ = i[1]
	return
}

type formula struct {
	outputs int
	terms   []term
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
	return m
}
func (f formulae) react(typ string, prs ...map[string]int) (pr map[string]int) {
	if len(prs) <= 0 {
		pr = make(map[string]int)
	} else {
		pr = prs[0]
	}
	pr[typ] += f[typ].outputs
	for _, trm := range f[typ].terms {
		if trm.typ == Ore {
			pr[Ore] += trm.n
			continue
		}
		for pr[trm.typ] < trm.n {
			pr = f.react(trm.typ, pr)
		}
		pr[trm.typ] -= trm.n
	}
	return
}

func part1(input [][]string) (rc int) {
	m := mkformulae(input)
	rc = m.react(Fuel)[Ore]
	return
}

func part2(input [][]string) int {
	return 0
}
