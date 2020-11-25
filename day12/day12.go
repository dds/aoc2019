package main

import (
	"fmt"

	"github.com/dds/aoc2019/lib"
)

var Input = lib.InputInts(lib.Inputs[12], lib.NumberParser)

func main() {
	fmt.Println(part1(Input))
	Input = lib.InputInts(lib.Inputs[12], lib.NumberParser)
	fmt.Println(part2(Input))
}

const N = 4

func part1(input [][]int) int {
	var p, v [N]column
	for i := 0; i < N; i++ {
		p[i] = column(input[i])
		v[i] = column([]int{0, 0, 0})
	}
	for i := 0; i < 1000; i++ {
		for dim := 0; dim < 3; dim++ {
			step(p, v, dim)
		}
	}
	return energy(p, v)
}

func gravity(p, v [N]column, dim int) {
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			if p[i][dim] == p[j][dim] {
				continue
			}
			if p[i][dim] < p[j][dim] {
				v[i][dim]++
				v[j][dim]--
			} else {
				v[i][dim]--
				v[j][dim]++
			}
		}
	}
}

func velocity(p, v [N]column, dim int) {
	for i := 0; i < N; i++ {
		p[i][dim] += v[i][dim]
	}
}

func step(p, v [N]column, dim int) {
	gravity(p, v, dim)
	velocity(p, v, dim)
}

func energy(p, v [N]column) (r int) {
	for i := 0; i < N; i++ {
		var pot, kin int
		for j := 0; j < 3; j++ {
			if p[i][j] < 0 {
				pot += p[i][j]
			} else {
				pot += -1 * p[i][j]
			}
			if v[i][j] < 0 {
				kin += v[i][j]
			} else {
				kin += -1 * v[i][j]
			}
		}
		r += pot * kin
	}
	return
}

func cycle(p, v [N]column, dim int) int {
	type s struct {
		p0, p1, p2, p3, v0, v1, v2, v3 int
	}
	m := map[s]int{}
	for i := 0; ; i++ {
		s := s{
			p0: p[0][dim], p1: p[1][dim], p2: p[2][dim], p3: p[3][dim],
			v0: v[0][dim], v1: v[1][dim], v2: v[2][dim], v3: v[3][dim],
		}
		if m[s] == 1 {
			return i
		}
		m[s] = 1
		step(p, v, dim)
	}
}

func part2(input [][]int) int {
	var p, v [N]column
	for i := 0; i < N; i++ {
		p[i] = column(input[i])
		v[i] = column([]int{0, 0, 0})
	}

	x := cycle(p, v, 0)
	fmt.Println("x cycle: ", x)

	y := cycle(p, v, 1)
	fmt.Println("y cycle: ", y)

	z := cycle(p, v, 2)
	fmt.Println("z cycle: ", z)
	return lcm(lcm(x, y), z)
}

func gcd(a, b int) int {
	max := a
	min := b
	if min > max {
		max, min = min, max
	}
	for {
		r := max % min
		if r == 0 {
			return min
		}
		max = min
		min = r
	}
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

type column []int
