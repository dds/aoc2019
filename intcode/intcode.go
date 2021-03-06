package intcode

import (
	"context"
	"fmt"
	"strconv"
)

type Code []int

type Verb int

const (
	Unknown Verb = iota
	Add
	Mul
	Input
	Output
	JumpIfTrue
	JumpIfFalse
	LessThan
	Equals
	SetBase
	Halt = 99
)

type Opmode int

func ParseVerb(c int) Verb {
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
	return Verb(i)
}

const (
	PositionMode  Opmode = 0
	ImmediateMode Opmode = 1
	RelativeMode  Opmode = 2
)

// Returns the parameter modes of an instruction.
func Opmodes(code int) (r []Opmode) {
	s := fmt.Sprint(code)
	n := len(s)
	switch op := ParseVerb(code); op {
	default:
		r = []Opmode{PositionMode}
	case Halt:
		r = []Opmode{}

	// Verbs of a single parameter.
	case SetBase:
		r = []Opmode{PositionMode}

	// 	Verbs that take two parameters
	case JumpIfTrue:
		fallthrough
	case JumpIfFalse:
		r = []Opmode{PositionMode, PositionMode}

		// Verbs that take three parameters
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
		switch s[i] {
		case '0':
			r[j] = PositionMode
		case '1':
			r[j] = ImmediateMode
		case '2':
			r[j] = RelativeMode
		}
		j++
	}
	return
}

// Exec ...
func (code Code) Exec(ctx context.Context, in <-chan int, out chan<- int) (err error) {
	defer close(out)
	ip := 0
	base := 0
	count := 0
	c := make([]int, len(code))
	copy(c, code)
	page := func(addr int) {
		if addr < len(c) {
			return
		}
		cc := make([]int, addr+1)
		copy(cc, c)
		c = cc
	}
	args := func(op Verb, modes []Opmode) (r []int) {
		for i, o := range modes {
			var addr int
			switch o {
			case PositionMode:
				// Dereference at position ip+i+1
				addr = c[ip+i+1]
			case ImmediateMode:
				// Return the value at ip+i+1
				addr = ip + i + 1
			case RelativeMode:
				// Dereference at position base+ip+i+1
				// fmt.Println("base", base, "c[ip+i+1]", c[ip+i+1])
				addr = base + c[ip+i+1]
			}
			if addr < 0 {
				panic(fmt.Errorf("invalid memory address: %v", addr))
			}
			page(addr)
			r = append(r, addr)
		}
		return
	}

	var op Verb
	for i := c[ip]; ; i = c[ip] {
		count++
		switch op = ParseVerb(i); op {
		default:
			err = fmt.Errorf("Unknown op: %v, ip: %v, code: %v", op, ip, c[ip])
			return
		case Halt:
			return
		case Add:
			a := args(op, Opmodes(i))
			c[a[2]] = c[a[0]] + c[a[1]]
			ip += 4
		case Mul:
			a := args(op, Opmodes(i))
			c[a[2]] = c[a[0]] * c[a[1]]
			ip += 4
		case Input:
			a := args(op, Opmodes(i))
			select {
			case <-ctx.Done():
				return ctx.Err()
			case u, ok := <-in:
				if !ok {
					return
				}
				// if Opmodes(i)[0] == RelativeMode {
				// 	fmt.Println(c[:ip], "-->", c[ip:])
				// 	fmt.Println("INPUT a[0]", a[0], "base", base, "base+c[ip+1]", base+c[ip+1])
				// }
				c[a[0]] = u
				ip += 2
				// if Opmodes(i)[0] == RelativeMode {
				// 	fmt.Println(c[:ip], "-->", c[ip:])
				// }
			}
		case Output:
			a := args(op, Opmodes(i))
			select {
			case <-ctx.Done():
				return ctx.Err()
			case out <- c[a[0]]:
			}
			ip += 2
		case JumpIfTrue:
			a := args(op, Opmodes(i))
			if c[a[0]] != 0 {
				ip = c[a[1]]
			} else {
				ip += 3
			}
		case JumpIfFalse:
			a := args(op, Opmodes(i))
			if c[a[0]] == 0 {
				ip = c[a[1]]
			} else {
				ip += 3
			}
		case LessThan:
			a := args(op, Opmodes(i))
			if c[a[0]] < c[a[1]] {
				c[a[2]] = 1
			} else {
				c[a[2]] = 0
			}
			ip += 4
		case Equals:
			a := args(op, Opmodes(i))
			if c[a[0]] == c[a[1]] {
				c[a[2]] = 1
			} else {
				c[a[2]] = 0
			}
			ip += 4
		case SetBase:
			a := args(op, Opmodes(i))
			// if Opmodes(i)[0] == RelativeMode {
			// 	fmt.Println("SETBASE a[0]", a[0], "base", base, "c[a[0]]", c[a[0]], "new base", base+c[a[0]])
			// }
			base += c[a[0]]
			// fmt.Println("base", base, "len(c)", len(c), "a", a, "ip", ip)
			// fmt.Println("c", c)
			ip += 2
		}
	}
}
