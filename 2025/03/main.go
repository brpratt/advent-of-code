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
	var (
		banks   = make([]bank, 0)
		scanner = bufio.NewScanner(r)
	)

	for scanner.Scan() {
		line := scanner.Text()
		banks = append(banks, NewBank(line))
	}

	return banks
}

func maxJoltage(b bank, num int) int {
	var (
		batteries = b.Batteries()
		idxs      = make([]int, num)
	)

	for i := range idxs {
		idx := 0

		if i > 0 {
			idx = idxs[i-1] + 1
		}

		for j := idx + 1; j < len(batteries)-(num-i-1); j++ {
			if batteries[j] > batteries[idx] {
				idx = j
			}
		}

		idxs[i] = idx
	}

	joltage := batteries[idxs[0]]
	for i := 1; i < len(idxs); i++ {
		joltage = (joltage * 10) + batteries[idxs[i]]
	}

	return joltage
}

func part01(banks []bank) int {
	var sum int

	for _, bank := range banks {
		sum += maxJoltage(bank, 2)
	}

	return sum
}

func part02(banks []bank) int {
	var sum int

	for _, bank := range banks {
		sum += maxJoltage(bank, 12)
	}

	return sum
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	banks := parseBanks(input)
	fmt.Println(part01(banks))
	fmt.Println(part02(banks))
}
