package y2023

import (
	"bufio"
	"fmt"
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

	// number is within the y window

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

func (sch schematic) show() {
	fmt.Printf("bound: (%2d, %2d)\n", sch.bound.x, sch.bound.y)

	fmt.Println("\nnumbers:")
	for _, num := range sch.numbers {
		fmt.Printf("\t%4d (%2d, %2d)\n", num.val, num.loc.x, num.loc.y)
	}
	fmt.Println("\nsymbols:")
	for _, sym := range sch.symbols {
		fmt.Printf("\t%s (%2d, %2d)\n", string(sym.val), sym.loc.x, sym.loc.y)
	}
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
	sch := parseSchematic(r)

	/*
		fmt.Println("\nadjacencies:")
		for isym, sym := range sch.symbols {
			for inum, num := range sch.numbers {
				fmt.Printf("\tsymbol %02d (%s) <=> number %02d (%4d): %v\n", isym, string(sym.val), inum, num.val, adjacent(sym, num))
			}
			fmt.Println()
		}
	*/

	var sum int

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
	sch := parseSchematic(r)

	var sum int

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

	return "", nil
}
