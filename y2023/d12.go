package y2023

import (
	"bufio"
	"io"
	"slices"
	"strconv"
	"strings"
)

func arrangements(row []rune, groups []int) int {
	if len(groups) == 0 {
		if slices.Contains(row, '#') {
			return 0
		}
		return 1
	}

	var count int

	for start := 0; start < len(row); start++ {
		if row[start] == '.' {
			continue
		}

		var length int

		for length = 1; length < groups[0]; length++ {
			if start+length == len(row) {
				break
			}

			if row[start+length] == '.' {
				break
			}
		}

		if length == groups[0] {
			if start+length < len(row) && row[start+length] != '#' {
				count += arrangements(row[start+length+1:], groups[1:])
			} else if start+length == len(row) {
				count += arrangements(row[start+length:], groups[1:])
			}
		}

		if row[start] == '#' {
			break
		}
	}

	return count
}

type springField struct {
	condition []rune
	groups    []int
}

func parseSpringFields(r io.Reader) (fields []springField) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fs := strings.Fields(scanner.Text())
		digits := strings.Split(fs[1], ",")

		field := springField{
			condition: []rune(fs[0]),
			groups:    make([]int, len(digits)),
		}

		for i, digit := range digits {
			field.groups[i], _ = strconv.Atoi(digit)
		}

		fields = append(fields, field)
	}
	return
}

func unfold(field springField) springField {
	newField := springField{
		condition: make([]rune, len(field.condition)*5+4),
		groups:    make([]int, len(field.groups)*5),
	}

	var ci int
	var gi int

	for i := 0; i < 5; i++ {
		for _, c := range field.condition {
			newField.condition[ci] = c
			ci++
		}

		if i != 4 {
			newField.condition[ci] = '?'
			ci++
		}

		for _, g := range field.groups {
			newField.groups[gi] = g
			gi++
		}
	}

	return newField
}

func SolveD12P01(r io.Reader) (string, error) {
	fields := parseSpringFields(r)

	var count int
	for _, field := range fields {
		n := arrangements(field.condition, field.groups)
		count += n
	}

	return strconv.Itoa(count), nil
}

func SolveD12P02(r io.Reader) (string, error) {
	fields := parseSpringFields(r)

	var count int
	for _, field := range fields {
		unfolded := unfold(field)
		n := arrangements(unfolded.condition, unfolded.groups)
		count += n
	}

	return strconv.Itoa(count), nil
}
