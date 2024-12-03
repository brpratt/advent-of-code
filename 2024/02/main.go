package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/brpratt/advent-of-code/file"
	"github.com/brpratt/advent-of-code/parse"
)

type report = []int

func parseReports(r io.Reader) []report {
	var reports []report

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var report report

		fields := strings.Fields(scanner.Text())
		for _, field := range fields {
			report = append(report, parse.MustAtoi(field))
		}

		reports = append(reports, report)
	}

	return reports
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func isDifferentSign(a, b int) bool {
	return (a < 0 && b >= 0) || (a >= 0 && b < 0)
}

func isWithinRange(n int) bool {
	absn := abs(n)
	return absn >= 1 && absn <= 3
}

func badIndex(report report) int {
	var prevDiff int
	var currDiff int

	prevDiff = report[1] - report[0]
	if !isWithinRange(prevDiff) {
		return 0
	}

	for i := 1; i < len(report)-1; i++ {
		currDiff = report[i+1] - report[i]
		if !isWithinRange(currDiff) {
			return i
		}

		if isDifferentSign(prevDiff, currDiff) {
			return i
		}

		prevDiff = currDiff
	}

	return -1
}

func part01(reports []report) int {
	var sum int

	for _, level := range reports {
		if badIndex(level) < 0 {
			sum++
		}
	}

	return sum
}

func part02(reports []report) int {
	var sum int

	for _, level := range reports {
		bad := badIndex(level)
		if bad < 0 {
			sum++
			continue
		}

		dampened := slices.Delete(slices.Clone(level), bad, bad+1)
		if badIndex(dampened) < 0 {
			sum++
			continue
		}

		dampened = slices.Delete(slices.Clone(level), bad+1, bad+2)
		if badIndex(dampened) < 0 {
			sum++
			continue
		}

		if bad == 1 {
			dampened = slices.Delete(slices.Clone(level), 0, 1)
			if badIndex(dampened) < 0 {
				sum++
			}
		}
	}

	return sum
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	levels := parseReports(input)
	fmt.Println(part01(levels))
	fmt.Println(part02(levels))
}
