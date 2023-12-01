package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"strings"

	aoc "github.com/brpratt/advent-of-code"
)

func usage() string {
	return fmt.Sprintf("Usage: %s <year> <day> <part>", os.Args[0])
}

func inputFile(year, day int) (*os.File, error) {
	file, err := os.Open(fmt.Sprintf("./y%d/input/d%02d.txt", year, day))
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, usage())
		os.Exit(1)
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, usage())
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, usage())
		os.Exit(1)
	}

	part, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Fprintln(os.Stderr, usage())
		os.Exit(1)
	}

	solver := aoc.Get(year, day, part)

	var input io.Reader

	inputFile, err := inputFile(year, day)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		input = strings.NewReader("")
	case err == nil:
		input = inputFile
		defer inputFile.Close()
	case err != nil:
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result, err := solver.Solve(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(result)
}
