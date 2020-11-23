package intcode

import (
	"fmt"
	"strconv"
)

type Opcode int

const (
	Unknown Opcode = iota
	Add
	Mul
	Input
	Output
	JumpIfTrue
	JumpIfFalse
	LessThan
	Equals
	Halt = 99
)

type Opmode int

func ParseOpcode(c int) Opcode {
	s := fmt.Sprint(c)
	var (
		i   int
		err error
	)
	if len(s) < 3 {
		i, err = strconv.Atoi(s)
	} else {
		i, err = strconv.Atoi(s[len(s)-2:])
	}
	if err != nil {
		return Unknown
	}
	return Opcode(i)
}

const (
	PositionMode  Opmode = 0
	ImmediateMode Opmode = 1
)

// Returns the parameter modes of an instruction.
func Opmodes(code int) (r []Opmode) {
	s := fmt.Sprint(code)
	n := len(s)
	switch op := ParseOpcode(code); op {
	default:
		r = []Opmode{PositionMode}
	case Halt:
		r = []Opmode{}

	// 	Opcodes that take two parameters
	case JumpIfTrue:
		fallthrough
	case JumpIfFalse:
		r = []Opmode{PositionMode, PositionMode}

		// Opcodes that take three parameters
	case LessThan:
		fallthrough
	case Equals:
		fallthrough
	case Add:
		fallthrough
	case Mul:
		r = []Opmode{PositionMode, PositionMode, PositionMode}
	}
	if len(s) < 3 {
		return
	}
	var j int
	for i := n - 3; i >= 0; i-- {
		if s[i] == '1' {
			r[j] = ImmediateMode
		}
		j++
	}
	return
}

// Exec ...
func Exec(code, in []int) (c, out []int, err error) {
	c = make([]int, len(code))
	copy(c, code)
	ip := 0

	args := func(op Opcode, modes []Opmode) (r []int) {
		for i, o := range modes {
			switch o {
			case PositionMode:
				// Dereference at position ip+i+1
				r = append(r, c[c[ip+i+1]])
			case ImmediateMode:
				// Return the value at ip+i+1
				r = append(r, c[ip+i+1])
			}
		}
		return
	}

	var op Opcode
	for i := c[ip]; ; i = c[ip] {
		switch op = ParseOpcode(i); op {
		default:
			err = fmt.Errorf("Unknown op: %v, ip: %v, code: %v", op, ip, c[ip])
			return
		case Halt:
			return
		case Add:
			a := args(op, Opmodes(i))
			c[c[ip+3]] = a[0] + a[1]
			ip += 4
		case Mul:
			a := args(op, Opmodes(i))
			c[c[ip+3]] = a[0] * a[1]
			ip += 4
		case Input:
			c[c[ip+1]] = in[0]
			in = in[1:]
			ip += 2
		case Output:
			// if Opmodes(i)[0] == ImmediateMode {
			// 	fmt.Println("WTF", Opmodes(i))
			// }
			// fmt.Println(c[ip-10 : ip+2])
			// fmt.Println(code[ip-10 : ip+2])
			a := args(op, Opmodes(i))
			out = append(out, a[0])
			ip += 2
		case JumpIfTrue:
			a := args(op, Opmodes(i))
			if a[0] != 0 {
				ip = a[1]
			} else {
				ip += 3
			}
		case JumpIfFalse:
			a := args(op, Opmodes(i))
			if a[0] == 0 {
				ip = a[1]
			} else {
				ip += 3
			}
		case LessThan:
			a := args(op, Opmodes(i))
			if a[0] < a[1] {
				c[c[ip+3]] = 1
			} else {
				c[c[ip+3]] = 0
			}
			ip += 4
		case Equals:
			a := args(op, Opmodes(i))
			if a[0] == a[1] {
				c[c[ip+3]] = 1
			} else {
				c[c[ip+3]] = 0
			}
			ip += 4
		}
	}
}
