package main

import "fmt"

func digits(number int) []int {
	ds := make([]int, 0, 6)

	for number > 0 {
		ds = append(ds, number%10)
		number /= 10
	}

	for i := 0; i < len(ds)/2; i++ {
		ds[i], ds[len(ds)-i-1] = ds[len(ds)-i-1], ds[i]
	}

	return ds
}

func neverDecreasing(number int) bool {
	ds := digits(number)

	for i := 0; i < len(ds)-1; i++ {
		if ds[i] > ds[i+1] {
			return false
		}
	}

	return true
}

func hasDouble(number int) bool {
	ds := digits(number)

	for i := 0; i < len(ds)-1; i++ {
		if ds[i] == ds[i+1] {
			return true
		}
	}

	return false
}

func hasBoundedDouble(number int) bool {
	ds := digits(number)

	for i := 0; i < len(ds)-1; {
		j := i + 1
		for j < len(ds) && ds[j] == ds[i] {
			j++
		}

		if j-i == 2 {
			return true
		}

		i = j
	}

	return false
}

func part01(first, last int) int {
	count := 0

	for i := first; i <= last; i++ {
		if hasDouble(i) && neverDecreasing(i) {
			count++
		}
	}

	return count
}

func part02(first, last int) int {
	count := 0

	for i := first; i <= last; i++ {
		if hasBoundedDouble(i) && neverDecreasing(i) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(part01(146810, 612564))
	fmt.Println(part02(146810, 612564))
}
