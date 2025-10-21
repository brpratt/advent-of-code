package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brpratt/advent-of-code/2019/intcode"
	"github.com/brpratt/advent-of-code/file"
)

func toIntcode(line string) []int {
	values := strings.Split(line, ",")
	program := make([]int, len(values))
	for i, v := range values {
		num, _ := strconv.Atoi(v)
		program[i] = num
	}

	return program
}

func runSerial(program []int, phases []int) int {
	channels := make([]chan int, len(phases)+1)

	for i := range channels {
		channels[i] = make(chan int)
	}

	for i, phase := range phases {
		c := intcode.NewComputer(program, channels[i], channels[i+1])
		go c.Run()
		channels[i] <- phase
	}

	channels[0] <- 0
	return <-channels[len(channels)-1]
}

func runFeedback(program []int, phases []int) int {
	channels := make([]chan int, len(phases)+1)

	for i := range channels {
		channels[i] = make(chan int, 1)
	}

	for i, phase := range phases {
		c := intcode.NewComputer(program, channels[i], channels[i+1])
		go c.Run()
		channels[i] <- phase
	}

	channels[0] <- 0
	var result int

	for {
		r, more := <-channels[len(channels)-1]
		if !more {
			break
		}
		result = r
		channels[0] <- r
	}

	return result
}

func heap(k int, nums []int, process func([]int)) {
	if k == 1 {
		process(nums)
	} else {
		heap(k-1, nums, process)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				nums[i], nums[k-1] = nums[k-1], nums[i]
			} else {
				nums[0], nums[k-1] = nums[k-1], nums[0]
			}

			heap(k-1, nums, process)
		}
	}
}

func permutations(nums []int) [][]int {
	perms := make([][]int, 0)

	add := func(p []int) {
		c := make([]int, len(p))
		copy(c, p)
		perms = append(perms, c)
	}

	heap(len(nums), nums, add)

	return perms
}

func part01(program []int) int {
	options := permutations([]int{0, 1, 2, 3, 4})
	max := runSerial(program, options[0])

	for i := 1; i < len(options); i++ {
		out := runSerial(program, options[i])
		if out > max {
			max = out
		}
	}

	return max
}

func part02(program []int) int {
	options := permutations([]int{5, 6, 7, 8, 9})
	max := runFeedback(program, options[0])

	for i := 1; i < len(options); i++ {
		out := runFeedback(program, options[i])
		if out > max {
			max = out
		}
	}

	return max
}

func main() {
	lines := file.Must(file.ReadLines("input.txt"))
	program := toIntcode(lines[0])

	fmt.Println(part01(program))
	fmt.Println(part02(program))
}
