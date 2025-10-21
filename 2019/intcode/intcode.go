package intcode

import (
	"fmt"
	"strings"

	"github.com/brpratt/advent-of-code/parse"
)

// Computer contains the execution of an intcode program
type Computer struct {
	pc     int
	rel    int
	in     <-chan int
	out    chan<- int
	signal chan<- int
	Memory []int
	Halted bool
}

// Signals
const (
	SignalIn = iota
	SignalOut
)

// Operations
const (
	OpAdd         = 1
	OpMultiply    = 2
	OpInput       = 3
	OpOutput      = 4
	OpJumpIfTrue  = 5
	OpJumpIfFalse = 6
	OpLessThan    = 7
	OpEquals      = 8
	OpAdjustRel   = 9
	OpHalt        = 99
)

// Parmeter modes
const (
	ModePosition  = 0
	ModeImmediate = 1
	ModeRelative  = 2
)

// Instruction represents the information contained in an opcode
type Instruction struct {
	Op    int
	Mode1 int
	Mode2 int
	Mode3 int
}

// Decode decodes an opcode into an Instruction
func Decode(opcode int) (Instruction, error) {
	op := opcode % 100

	if op < 1 {
		return Instruction{}, fmt.Errorf("illegal operation %d in opcode %d", op, opcode)
	}

	if op > 9 && op != 99 {
		return Instruction{}, fmt.Errorf("illegal operation %d in opcode %d", op, opcode)
	}

	mode1 := (opcode / 100) % 10
	if mode1 != ModePosition && mode1 != ModeImmediate && mode1 != ModeRelative {
		return Instruction{}, fmt.Errorf("illegal mode %d for parameter 1 in opcode %d", mode1, opcode)
	}

	mode2 := (opcode / 1000) % 10
	if mode2 != ModePosition && mode2 != ModeImmediate && mode2 != ModeRelative {
		return Instruction{}, fmt.Errorf("illegal mode %d for parameter 2 in opcode %d", mode2, opcode)
	}

	mode3 := (opcode / 10000) % 10
	if mode3 != ModePosition && mode3 != ModeImmediate && mode3 != ModeRelative {
		return Instruction{}, fmt.Errorf("illegal mode %d for parameter 3 in opcode %d", mode3, opcode)
	}

	return Instruction{op, mode1, mode2, mode3}, nil
}

func (c *Computer) expand(addr int) {
	if addr >= len(c.Memory) {
		memory := make([]int, addr+1)
		copy(memory, c.Memory)
		c.Memory = memory
	}
}

func (c *Computer) read(addr int) int {
	c.expand(addr)
	return c.Memory[addr]
}

func (c *Computer) write(addr, value int) {
	c.expand(addr)
	c.Memory[addr] = value
}

func (c *Computer) fetch(param, mode int) int {
	if mode == ModePosition {
		return c.read(param)
	}

	if mode == ModeRelative {
		return c.read(param + c.rel)
	}

	return param
}

func (c *Computer) place(param, value, mode int) {
	if mode == ModePosition {
		c.write(param, value)
	} else {
		c.write(param+c.rel, value)
	}
}

func (c *Computer) step() error {
	op, err := Decode(c.Memory[c.pc])
	if err != nil {
		return fmt.Errorf("failed decode at address %d: %v", c.pc, err)
	}

	switch op.Op {
	case OpAdd:
		x := c.fetch(c.Memory[c.pc+1], op.Mode1)
		y := c.fetch(c.Memory[c.pc+2], op.Mode2)
		c.place(c.Memory[c.pc+3], x+y, op.Mode3)
		c.pc += 4
	case OpMultiply:
		x := c.fetch(c.Memory[c.pc+1], op.Mode1)
		y := c.fetch(c.Memory[c.pc+2], op.Mode2)
		c.place(c.Memory[c.pc+3], x*y, op.Mode3)
		c.pc += 4
	case OpInput:
		if c.signal != nil {
			c.signal <- SignalIn
		}
		c.place(c.Memory[c.pc+1], <-c.in, op.Mode1)
		c.pc += 2
	case OpOutput:
		if c.signal != nil {
			c.signal <- SignalOut
		}
		c.out <- c.fetch(c.Memory[c.pc+1], op.Mode1)
		c.pc += 2
	case OpJumpIfTrue:
		x := c.fetch(c.Memory[c.pc+1], op.Mode1)
		y := c.fetch(c.Memory[c.pc+2], op.Mode2)

		if x == 0 {
			c.pc += 3
		} else {
			c.pc = y
		}
	case OpJumpIfFalse:
		x := c.fetch(c.Memory[c.pc+1], op.Mode1)
		y := c.fetch(c.Memory[c.pc+2], op.Mode2)

		if x == 0 {
			c.pc = y
		} else {
			c.pc += 3
		}
	case OpLessThan:
		x := c.fetch(c.Memory[c.pc+1], op.Mode1)
		y := c.fetch(c.Memory[c.pc+2], op.Mode2)

		if x < y {
			c.place(c.Memory[c.pc+3], 1, op.Mode3)
		} else {
			c.place(c.Memory[c.pc+3], 0, op.Mode3)
		}

		c.pc += 4
	case OpEquals:
		x := c.fetch(c.Memory[c.pc+1], op.Mode1)
		y := c.fetch(c.Memory[c.pc+2], op.Mode2)

		if x == y {
			c.place(c.Memory[c.pc+3], 1, op.Mode3)
		} else {
			c.place(c.Memory[c.pc+3], 0, op.Mode3)
		}

		c.pc += 4
	case OpAdjustRel:
		x := c.fetch(c.Memory[c.pc+1], op.Mode1)
		c.rel += x
		c.pc += 2
	case OpHalt:
		c.Halted = true
		if c.signal != nil {
			close(c.signal)
		}
		close(c.out)
	}

	return nil
}

// NewComputer returns a new intcode computer intialized from program
func NewComputer(program []int, in <-chan int, out chan<- int) Computer {
	c := Computer{
		Memory: make([]int, len(program)),
		in:     in,
		out:    out,
	}

	copy(c.Memory, program)

	return c
}

func (c *Computer) AddSignalHandler(ch chan int) {
	c.signal = ch
}

// Run runs the program until halted
func (c *Computer) Run() error {
	for !c.Halted {
		if err := c.step(); err != nil {
			return err
		}
	}

	return nil
}

// FromText parses the intcode text into a program
//
// In text form, intcode programs are comma-separated numbers
func FromText(text string) []int {
	values := strings.Split(text, ",")
	program := make([]int, len(values))

	for i, v := range values {
		program[i] = parse.MustAtoi(v)
	}

	return program
}
