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
			// "Parameters that an instruction writes to will never be in immediate mode."
			if i == 2 && (op == Add || op == Mul) {
				r = append(r, c[ip+3])
				return
			}
			switch o {
			case PositionMode:
				// Return the value at position ip+i+1
				r = append(r, c[ip+i+1])
			case ImmediateMode:
				// Return ip+i+1
				r = append(r, ip+i+1)
			}
		}
		return
	}

	var op Opcode
	for i := c[ip]; ; i = c[ip] {
		a := args(op, Opmodes(i))
		switch op = ParseOpcode(i); op {
		default:
			err = fmt.Errorf("Unknown op: %v, ip: %v, code: %v", op, ip, c[ip])
			return
		case Halt:
			return
		case Add:
			c[a[2]] = c[a[0]] + c[a[1]]
			ip += 4
		case Mul:
			c[a[2]] = c[a[0]] * c[a[1]]
			ip += 4
		case Input:
			c[a[0]] = in[0]
			in = in[1:]
			ip += 2
		case Output:
			fmt.Println(c[ip-10 : ip+2])
			out = append(out, c[c[ip+1]])
			ip += 2
		case JumpIfTrue:
			if c[a[0]] != 0 {
				ip = c[a[1]]
			}
			ip += 2
		case JumpIfFalse:
			if c[a[0]] == 0 {
				ip = c[a[1]]
			}
			ip += 2
		case LessThan:
			if c[a[0]] < c[a[1]] {
				c[a[2]] = 1
			} else {
				c[a[2]] = 0
			}
			ip += 3
		case Equals:
			if c[a[0]] == c[a[1]] {
				c[a[2]] = 1
			} else {
				c[a[2]] = 0
			}
			ip += 3
		}
	}
}
