package intcode

import (
	"reflect"
	"testing"
)

func TestNewComputer(t *testing.T) {
	programs := [][]int{
		{1, 2, 3, 4},
		{0},
	}

	for _, program := range programs {
		c := NewComputer(program, make(chan int), make(chan int))

		if !reflect.DeepEqual(program, c.Memory) {
			t.Errorf("expected memory %v, got %v", program, c.Memory)
		}

		c.Memory[0] = c.Memory[0] + 1
		if program[0] == c.Memory[0] {
			t.Error("expected initializing program and mem to be separate slices")
		}
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		opcode      int
		shouldError bool
		expected    Instruction
	}{
		{1, false, Instruction{OpAdd, ModePosition, ModePosition, ModePosition}},
		{2, false, Instruction{OpMultiply, ModePosition, ModePosition, ModePosition}},
		{3, false, Instruction{OpInput, ModePosition, ModePosition, ModePosition}},
		{4, false, Instruction{OpOutput, ModePosition, ModePosition, ModePosition}},
		{99, false, Instruction{OpHalt, ModePosition, ModePosition, ModePosition}},
		{17, true, Instruction{}},
		{-1, true, Instruction{}},
		{1002, false, Instruction{OpMultiply, ModePosition, ModeImmediate, ModePosition}},
		{11101, false, Instruction{OpAdd, ModeImmediate, ModeImmediate, ModeImmediate}},
		{20001, false, Instruction{OpAdd, ModePosition, ModePosition, ModeRelative}},
	}

	for _, test := range tests {
		instr, err := Decode(test.opcode)

		if test.shouldError && err == nil {
			t.Errorf("opcode %d -- expected error, got %v", test.opcode, instr)
			continue
		}

		if !test.shouldError && err != nil {
			t.Errorf("opcode %d -- unexpected error %v", test.opcode, err)
			continue
		}

		if !test.shouldError && !reflect.DeepEqual(test.expected, instr) {
			t.Errorf("opcode %d -- expected instruction %v, got %v", test.opcode, test.expected, instr)
		}
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		program []int
		final   []int
		input   []int
		output  []int
	}{
		{
			"Add 1",
			[]int{1, 0, 0, 0, 99},
			[]int{2, 0, 0, 0, 99},
			[]int{},
			[]int{},
		},
		{
			"Multiply 1",
			[]int{2, 4, 4, 5, 99, 0},
			[]int{2, 4, 4, 5, 99, 9801},
			[]int{},
			[]int{},
		},
		{
			"Add and multiply 1",
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50},
			[]int{},
			[]int{},
		},
		{
			"Add and multiply 2",
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			[]int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			[]int{},
			[]int{},
		},
		{
			"Input single value",
			[]int{3, 3, 99, 0},
			[]int{3, 3, 99, 15},
			[]int{15},
			[]int{},
		},
		{
			"Input multiple values",
			[]int{3, 6, 3, 0, 99, 0, 0},
			[]int{44, 6, 3, 0, 99, 0, 2},
			[]int{2, 44},
			[]int{},
		},
		{
			"Output single value",
			[]int{4, 0, 99},
			[]int{4, 0, 99},
			[]int{},
			[]int{4},
		},
		{
			"Multiply immediate mode",
			[]int{1002, 4, 3, 4, 33},
			[]int{1002, 4, 3, 4, 99},
			[]int{},
			[]int{},
		},
	}

	for _, test := range tests {
		in := make(chan int)
		out := make(chan int)
		inputDone := make(chan bool)
		outputDone := make(chan bool)
		var inputCount int
		output := make([]int, 0)

		c := NewComputer(test.program, in, out)

		go func() {
			for _, v := range test.input {
				inputCount++
				in <- v
			}

			close(in)
			inputDone <- true
		}()

		go func() {
			for {
				v, more := <-out
				if !more {
					outputDone <- true
					return
				}
				output = append(output, v)
			}
		}()

		err := c.Run()
		<-inputDone
		<-outputDone

		if err != nil {
			t.Errorf("unexpected error in program '%s': %v", test.name, err)
			continue
		}

		if !reflect.DeepEqual(test.final, c.Memory) {
			t.Errorf("unexpected memory layout in program '%s': expected %v, got %v", test.name, test.final, c.Memory)
			continue
		}

		if inputCount != len(test.input) {
			t.Errorf("failed to read all input for program '%s'", test.name)
			continue
		}

		if !reflect.DeepEqual(test.output, output) {
			t.Errorf("unexpected output for program '%s': expected %q, got %q", test.name, test.output, output)
		}
	}
}

func TestMoreOps(t *testing.T) {
	tests := []struct {
		name    string
		program []int
		input   []int
		output  []int
	}{
		{
			"Test equal to 8 (1)",
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{8},
			[]int{1},
		},
		{
			"Test equal to 8 (2)",
			[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{7},
			[]int{0},
		},
		{
			"Test less than 8 (1)",
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{8},
			[]int{0},
		},
		{
			"Test less than 8 (2)",
			[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			[]int{3},
			[]int{1},
		},
		{
			"Test not equal to 0 (1)",
			[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			[]int{0},
			[]int{0},
		},
		{
			"Test not equal to 0 (2)",
			[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			[]int{1},
			[]int{1},
		},
		{
			"Test not equal to 0 (3)",
			[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			[]int{0},
			[]int{0},
		},
		{
			"Test not equal to 0 (4)",
			[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			[]int{1},
			[]int{1},
		},
		{
			"Test less than, equal, or greater than 8 (1)",
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int{7},
			[]int{999},
		},
		{
			"Test less than, equal, or greater than 8 (1)",
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int{8},
			[]int{1000},
		},
		{
			"Test less than, equal, or greater than 8 (1)",
			[]int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			[]int{9},
			[]int{1001},
		},
	}

	for _, test := range tests {
		in := make(chan int)
		out := make(chan int)
		inputDone := make(chan bool)
		outputDone := make(chan bool)
		var inputCount int
		output := make([]int, 0)

		c := NewComputer(test.program, in, out)

		go func() {
			for _, v := range test.input {
				inputCount++
				in <- v
			}

			close(in)
			inputDone <- true
		}()

		go func() {
			for {
				v, more := <-out
				if !more {
					outputDone <- true
					return
				}
				output = append(output, v)
			}
		}()

		err := c.Run()
		if err != nil {
			t.Errorf("unexpected error in program '%s': %v", test.name, err)
			continue
		}

		<-inputDone
		<-outputDone

		if inputCount != len(test.input) {
			t.Errorf("failed to read all input for program '%s'", test.name)
			continue
		}

		if !reflect.DeepEqual(test.output, output) {
			t.Errorf("unexpected output for program '%s': expected %q, got %q", test.name, test.output, output)
		}
	}
}

func TestCompleteFeatures(t *testing.T) {
	tests := []struct {
		name    string
		program []int
		input   []int
		output  []int
	}{
		{
			"Test relative adjust",
			[]int{109, 2, 204, -2, 99},
			[]int{},
			[]int{109},
		},
		{
			"Test quine",
			[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			[]int{},
			[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			"Test output 16 digit number",
			[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			[]int{},
			[]int{1219070632396864},
		},
		{
			"Test output middle number",
			[]int{104, 1125899906842624, 99},
			[]int{},
			[]int{1125899906842624},
		},
	}

	for _, test := range tests {
		in := make(chan int)
		out := make(chan int)
		inputDone := make(chan bool)
		outputDone := make(chan bool)
		var inputCount int
		output := make([]int, 0)

		c := NewComputer(test.program, in, out)

		go func() {
			for _, v := range test.input {
				inputCount++
				in <- v
			}

			close(in)
			inputDone <- true
		}()

		go func() {
			for {
				v, more := <-out
				if !more {
					outputDone <- true
					return
				}
				output = append(output, v)
			}
		}()

		err := c.Run()
		if err != nil {
			t.Errorf("unexpected error in program '%s': %v", test.name, err)
			continue
		}

		<-inputDone
		<-outputDone

		if inputCount != len(test.input) {
			t.Errorf("failed to read all input for program '%s'", test.name)
			continue
		}

		if !reflect.DeepEqual(test.output, output) {
			t.Errorf("unexpected output for program '%s': expected %q, got %q", test.name, test.output, output)
		}
	}
}
