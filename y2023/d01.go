package y2023

import (
	"io"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/brpratt/advent-of-code/extract"
)

func extractNumber(line string) int {
	index := 0
	runes := []rune{0, 0}

	for _, letter := range line {
		if unicode.IsDigit(letter) {
			if index == 0 {
				runes[0] = letter
				index += 1
			}

			runes[1] = letter
		}
	}

	number, _ := strconv.Atoi(string(runes))

	return number
}

func numberPrefix(s string) (rune, bool) {
	prefixes := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	for k, v := range prefixes {
		if strings.HasPrefix(s, k) {
			return v, true
		}
	}

	firstRune, _ := utf8.DecodeRuneInString(s)
	if unicode.IsDigit(firstRune) {
		return firstRune, true
	}

	return utf8.RuneError, false
}

func extractNumber2(line string) int {
	index := 0
	runes := []rune{0, 0}
	lineRunes := []rune(line)

	for i := range lineRunes {
		if r, ok := numberPrefix(string(lineRunes[i:])); ok {
			if index == 0 {
				runes[0] = r
				index += 1
			}

			runes[1] = r
		}
	}

	number, _ := strconv.Atoi(string(runes))

	return number
}

func SolveD01P01(r io.Reader) (string, error) {
	sum := 0
	lines := extract.Lines(r)
	for _, line := range lines {
		number := extractNumber(line)
		sum += number
	}

	return strconv.Itoa(sum), nil
}

func SolveD01P02(r io.Reader) (string, error) {
	sum := 0
	lines := extract.Lines(r)
	for _, line := range lines {
		number := extractNumber2(line)
		sum += number
	}

	return strconv.Itoa(sum), nil
}
