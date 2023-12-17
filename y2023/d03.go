package y2023

import (
	"bufio"
	"io"
	"strconv"
	"unicode"

	"github.com/brpratt/advent-of-code/grid"
)

type number struct {
	loc    grid.Point
	val    int
	digits int
}

type symbol struct {
	loc grid.Point
	val rune
}

func adjacent(s symbol, n number) bool {
	if n.loc.Y < s.loc.Y-1 {
		return false
	}

	if n.loc.Y > s.loc.Y+1 {
		return false
	}

	if n.loc.X > s.loc.X+1 {
		return false
	}

	if n.loc.X < s.loc.X-n.digits {
		return false
	}

	return true
}

type schematic struct {
	bound   grid.Point
	numbers []number
	symbols []symbol
}

func parseSchematic(r io.Reader) (sch schematic) {
	var loc grid.Point
	var parsingNum bool
	var num number
	reader := bufio.NewReader(r)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}

		if parsingNum {
			if unicode.IsDigit(rune(b)) {
				val, _ := strconv.Atoi(string(b))
				num.val = (num.val * 10) + val
				num.digits += 1
				loc.X += 1
				continue
			} else {
				sch.numbers = append(sch.numbers, num)
				num.loc.X = 0
				num.loc.Y = 0
				num.val = 0
				num.digits = 0
				parsingNum = false
			}
		}

		switch {
		case b == '.':
			loc.X += 1
		case b == '\n':
			sch.bound.X = loc.X
			loc.X = 0
			loc.Y += 1
		case unicode.IsDigit(rune(b)):
			parsingNum = true
			val, _ := strconv.Atoi(string(b))
			num.val = (num.val * 10) + val
			num.digits += 1
			num.loc.X = loc.X
			num.loc.Y = loc.Y
			loc.X += 1
		default: // symbol
			sym := symbol{
				val: rune(b),
				loc: grid.Point{
					X: loc.X,
					Y: loc.Y,
				},
			}
			sch.symbols = append(sch.symbols, sym)
			loc.X += 1
		}
	}

	sch.bound.Y = loc.Y + 1

	return
}

func SolveD03P01(r io.Reader) (string, error) {
	var sum int
	sch := parseSchematic(r)

	for _, sym := range sch.symbols {
		for _, num := range sch.numbers {
			if adjacent(sym, num) {
				sum += num.val
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func SolveD03P02(r io.Reader) (string, error) {
	var sum int
	sch := parseSchematic(r)

	for _, sym := range sch.symbols {
		if sym.val != '*' {
			continue
		}

		var first int
		var firstFound bool

		for _, num := range sch.numbers {
			if !adjacent(sym, num) {
				continue
			}

			if !firstFound {
				first = num.val
				firstFound = true
				continue
			}

			sum += first * num.val

			first = 0
			firstFound = false

			break
		}
	}

	return strconv.Itoa(sum), nil
}
