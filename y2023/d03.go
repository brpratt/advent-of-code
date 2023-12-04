package y2023

import (
	"bufio"
	"io"
	"strconv"
	"unicode"
)

type point struct {
	x int
	y int
}

type number struct {
	loc    point
	val    int
	digits int
}

type symbol struct {
	loc point
	val rune
}

func adjacent(s symbol, n number) bool {
	if n.loc.y < s.loc.y-1 {
		return false
	}

	if n.loc.y > s.loc.y+1 {
		return false
	}

	if n.loc.x > s.loc.x+1 {
		return false
	}

	if n.loc.x < s.loc.x-n.digits {
		return false
	}

	return true
}

type schematic struct {
	bound   point
	numbers []number
	symbols []symbol
}

func parseSchematic(r io.Reader) (sch schematic) {
	var loc point
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
				loc.x += 1
				continue
			} else {
				sch.numbers = append(sch.numbers, num)
				num.loc.x = 0
				num.loc.y = 0
				num.val = 0
				num.digits = 0
				parsingNum = false
			}
		}

		switch {
		case b == '.':
			loc.x += 1
		case b == '\n':
			sch.bound.x = loc.x
			loc.x = 0
			loc.y += 1
		case unicode.IsDigit(rune(b)):
			parsingNum = true
			val, _ := strconv.Atoi(string(b))
			num.val = (num.val * 10) + val
			num.digits += 1
			num.loc.x = loc.x
			num.loc.y = loc.y
			loc.x += 1
		default: // symbol
			sym := symbol{
				val: rune(b),
				loc: point{
					x: loc.x,
					y: loc.y,
				},
			}
			sch.symbols = append(sch.symbols, sym)
			loc.x += 1
		}
	}

	sch.bound.y = loc.y + 1

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
