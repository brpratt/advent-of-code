package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/brpratt/advent-of-code/file"
)

func solve(input io.Reader) (string, string) {
	scanner := bufio.NewScanner(input)

	calories := []int{}
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			calories = append(calories, count)
			count = 0
			continue
		}

		calorie, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		count += calorie
	}

	calories = append(calories, count)
	sort.Ints(calories)

	max := calories[len(calories)-1]
	topThree := calories[len(calories)-1] + calories[len(calories)-2] + calories[len(calories)-3]

	return strconv.Itoa(max), strconv.Itoa(topThree)
}

func main() {
	input := file.Must(os.Open("input.txt"))
	defer input.Close()

	part1, part2 := solve(input)
	fmt.Println(part1)
	fmt.Println(part2)
}
