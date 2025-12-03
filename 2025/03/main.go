package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/brpratt/advent-of-code/file"
)

type bank []int

func NewBank(rep string) bank {
	b := make([]int, len(rep))
	for i, ch := range rep {
		b[i] = int(ch) - 0x30
	}
	return b
}

func (b bank) Batteries() []int {
	return []int(b)
}

func parseBanks(r io.Reader) []bank {
	banks := make([]bank, 0)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		banks = append(banks, NewBank(line))
	}

	return banks
}

func maxJoltage(b bank) int {
	batteries := b.Batteries()
	leftidx := 0

	for i := 1; i < len(batteries)-1; i++ {
		if batteries[i] > batteries[leftidx] {
			leftidx = i
		}
	}

	rightidx := leftidx + 1
	for i := rightidx + 1; i < len(batteries); i++ {
		if batteries[i] > batteries[rightidx] {
			rightidx = i
		}
	}

	return (batteries[leftidx] * 10) + batteries[rightidx]
}

func part01(banks []bank) int {
	var sum int

	for _, bank := range banks {
		sum += maxJoltage(bank)
	}

	return sum
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	banks := parseBanks(input)
	fmt.Println(part01(banks))
}
