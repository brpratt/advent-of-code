package extract

import (
	"bufio"
	"io"
)

func Lines(r io.Reader) []string {
	lines := make([]string, 0)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
