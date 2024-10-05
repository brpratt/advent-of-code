package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/brpratt/advent-of-code/file"
)

type history []int

type report []history

func parseReport(r io.Reader) report {
	rep := make([]history, 0)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fs := strings.Fields(scanner.Text())
		hist := make([]int, len(fs))
		for i := range fs {
			hist[i], _ = strconv.Atoi(fs[i])
		}
		rep = append(rep, hist)
	}

	return rep
}

type ehistory [][]int

func extend(h history) ehistory {
	eh := make(ehistory, 0)
	eh = append(eh, make([]int, len(h)))
	copy(eh[0], h)

	level := 0
	done := false
	for !done {
		done = true
		eh = append(eh, make([]int, len(eh[level])-1))
		for i := range eh[level+1] {
			eh[level+1][i] = eh[level][i+1] - eh[level][i]
			if eh[level+1][i] != 0 {
				done = false
			}
		}
		level++
	}

	return eh
}

func extrapr(eh ehistory) int {
	level := len(eh) - 1
	eh[level] = append(eh[level], 0)

	for level > 0 {
		level--
		i := len(eh[level+1])
		eh[level] = append(eh[level], eh[level][i-1]+eh[level+1][i-1])
	}

	return eh[0][len(eh[0])-1]
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	report := parseReport(input)

	fmt.Println(part01(report))
	fmt.Println(part02(report))
}

func part01(report report) int {
	var sum int
	for _, hist := range report {
		eh := extend(hist)
		sum += extrapr(eh)
	}

	return sum
}

func extrapl(eh ehistory) int {
	level := len(eh) - 1
	eh[level] = append([]int{0}, eh[level]...)

	for level > 0 {
		level--
		val := eh[level][0] - eh[level+1][0]
		eh[level] = append([]int{val}, eh[level]...)
	}

	return eh[0][0]
}

func part02(report report) int {
	var sum int
	for _, hist := range report {
		eh := extend(hist)
		sum += extrapl(eh)
	}

	return sum
}
