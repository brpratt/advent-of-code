package main

import (
	"fmt"

	"github.com/brpratt/advent-of-code/file"
)

func main() {
	lines := file.Must(file.ReadLines("input.txt"))

	solvePart01(lines)
}

func solvePart01(numbers []string) {
	numCount := len(numbers)
	bitCount := len(numbers[0])
	setBitCount := make([]int, bitCount)

	for _, number := range numbers {
		for bit := 0; bit < bitCount; bit++ {
			if number[bitCount-(bit+1)] == '1' {
				setBitCount[bit]++
			}
		}
	}

	var gamma int
	var epsilon int

	for bit := 0; bit < bitCount; bit++ {
		if setBitCount[bit] > numCount/2 {
			gamma |= (1 << bit)
		} else {
			epsilon |= (1 << bit)
		}
	}

	fmt.Println(gamma * epsilon)
}
