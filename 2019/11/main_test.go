package main

import (
	"testing"

	"github.com/brpratt/advent-of-code/2019/intcode"
	"github.com/brpratt/advent-of-code/file"
)

func TestPart01(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		program := []int{3, 0, 104, 1, 99}
		expected := 1
		got := part01(program)

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("actual", func(t *testing.T) {
		lines := file.Must(file.ReadLines("input.txt"))
		program := intcode.FromText(lines[0])

		expected := 1967
		got := part01(program)

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}

func TestPart02(t *testing.T) {
	lines := file.Must(file.ReadLines("input.txt"))
	program := intcode.FromText(lines[0])

	expected := ` #  # ###  #  # ####  ##  #### ###  #  #   
 # #  #  # #  # #    #  #    # #  # # #    
 ##   ###  #  # ###  #      #  ###  ##     
 # #  #  # #  # #    # ##  #   #  # # #    
 # #  #  # #  # #    #  # #    #  # # #    
 #  # ###   ##  ####  ### #### ###  #  #   
`

	got := part02(program)

	if got != expected {
		t.Errorf("expected:\n%s\n got:\n%s", expected, got)
	}
}
