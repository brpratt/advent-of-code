package y2023

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

func sum(nums []int) (s int) {
	for _, v := range nums {
		s += v
	}
	return
}

func slots(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return sum(nums) + len(nums) - 1
}

func hash(row []rune, groups []int) string {
	var sb strings.Builder

	fmt.Fprint(&sb, string(row))
	if len(groups) > 0 {
		for _, v := range groups[1:] {
			fmt.Fprintf(&sb, "-%d", v)
		}
	}

	return sb.String()
}

// we need to track _potential_ locations for each group
func arrangements(row []rune, groups []int) int {
	var cache = make(map[string]int)

	var iter func(_row []rune, _groups []int) int

	iter = func(_row []rune, _groups []int) int {
		h := hash(_row, _groups)
		if v, ok := cache[h]; ok {
			return v
		}

		// the row must be at least as long as:
		// - the total length of all the groups
		// - PLUS the space between the groups
		if len(_row) < slots(_groups) {
			return 0
		}

		if len(_groups) == 0 {
			if slices.Contains(_row, '#') {
				return 0
			}
			return 1
		}

		var count int
		end := len(_row) - slots(_groups[1:])

		for start := 0; start < end; start++ {
			if _row[start] == '.' {
				continue
			}

			var pos int

			// walk forward until we either:
			// - reach the desired length (loop condition)
			// - reach the end of the row (1st if)
			// - reach a '.' (2nd if)
			for pos = start + 1; pos-start < _groups[0]; pos++ {
				if pos == end {
					break
				}

				if _row[pos] == '.' {
					break
				}
			}

			// check if we're long enough to consider this position valid
			if pos-start == _groups[0] {
				if pos < end && _row[pos] != '#' {
					count += iter(_row[pos+1:], _groups[1:])
				} else if pos == end && len(_groups[1:]) == 0 {
					count += iter(_row[pos:], _groups[1:])
				}
			}

			if _row[start] == '#' {
				break
			}
		}

		cache[h] = count
		return count
	}

	return iter(row, groups)
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
