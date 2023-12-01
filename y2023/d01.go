package y2023

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func readLines(r io.Reader) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

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

func toNumber(runes []rune) (int, bool) {
	str := string(runes)
	switch {
	case strings.HasPrefix(str, "one"):
		return 1, true
	case strings.HasPrefix(str, "two"):
		return 2, true
	case strings.HasPrefix(str, "three"):
		return 3, true
	case strings.HasPrefix(str, "four"):
		return 4, true
	case strings.HasPrefix(str, "five"):
		return 5, true
	case strings.HasPrefix(str, "six"):
		return 6, true
	case strings.HasPrefix(str, "seven"):
		return 7, true
	case strings.HasPrefix(str, "eight"):
		return 8, true
	case strings.HasPrefix(str, "nine"):
		return 9, true
	case unicode.IsDigit(runes[0]):
		return int(runes[0] - '0'), true
	default:
		return 0, false
	}
}

func extractNumber2(line string) int {
	index := 0
	nums := []int{0, 0}
	runes := []rune(line)

	for i := range runes {
		if num, ok := toNumber(runes[i:]); ok {
			if index == 0 {
				nums[0] = num
				index += 1
			}

			nums[1] = num
		}
	}

	number, _ := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[1]))

	return number
}

func SolveD01P01(r io.Reader) (string, error) {
	sum := 0
	lines := readLines(r)
	for _, line := range lines {
		number := extractNumber(line)
		sum += number
	}

	return strconv.Itoa(sum), nil
}

func SolveD01P02(r io.Reader) (string, error) {
	sum := 0
	lines := readLines(r)
	for _, line := range lines {
		number := extractNumber2(line)
		sum += number
	}

	return strconv.Itoa(sum), nil
}
