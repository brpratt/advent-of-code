package main

import (
	"fmt"
	"strconv"

	"github.com/brpratt/advent-of-code/file"
)

func main() {
	lines := file.Must(file.ReadLines("input.txt"))

	depths := make([]int, len(lines))
	for i, v := range lines {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		depths[i] = num
	}

	solvePart01(depths)
	solvePart02(depths)
}

func solvePart01(depths []int) {
	prev := depths[0]
	ninc := 0
	for i := 1; i < len(depths); i++ {
		curr := depths[i]
		if curr-prev > 0 {
			ninc++
		}
		prev = curr
	}
	fmt.Println(ninc)
}

func solvePart02(depths []int) {
	prev := depths[0] + depths[1] + depths[2]
	ninc := 0
	for i := 3; i < len(depths); i++ {
		curr := depths[i-2] + depths[i-1] + depths[i]
		if curr-prev > 0 {
			ninc++
		}
		prev = curr
	}
	fmt.Println(ninc)
}
