package intcode

import "fmt"

type Opcode int

const (
	Unknown Opcode = iota
	Add
	Mul
	Input
	Output
	Halt = 99
)

// Exec ...
func Exec(code, in []int) (c, out []int, err error) {
	c = make([]int, len(code))
	copy(c, code)
	i := 0
	for op := Opcode(c[i]); op != Halt; op = Opcode(c[i]) {
		switch op {
		default:
			err = fmt.Errorf("Unknown op: %v, i: %v, code: %v", op, i, c[i])
			return
		case Add:
			args := [3]int{c[i+1], c[i+2], c[i+3]}
			c[args[2]] = c[args[0]] + c[args[1]]
			i += 4
		case Mul:
			args := [3]int{c[i+1], c[i+2], c[i+3]}
			c[args[2]] = c[args[0]] * c[args[1]]
			i += 4
		case Input:
			c[c[i+1]] = in[0]
			in = in[1:]
			i += 2
		case Output:
			out = append(out, c[c[i+1]])
			i += 2
		}
	}
	return
}
