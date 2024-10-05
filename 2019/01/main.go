package main

import (
	"fmt"

	"github.com/brpratt/advent-of-code/file"
)

func main() {
	masses := file.Must(file.ReadNumbers("input.txt"))

	fmt.Println(part01(masses))
	fmt.Println(part02(masses))
}

func part01(masses []int) int {
	fuel := 0

	for _, mass := range masses {
		fuel += calculateFuel(mass)
	}

	return fuel
}

func part02(masses []int) int {
	fuel := 0

	for _, mass := range masses {
		fuel += calculateFuelWithOverhead(mass)
	}

	return fuel
}

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateFuelWithOverhead(mass int) int {
	fuel := calculateFuel(mass)
	totalFuel := fuel

	for fuel > 0 {
		fuel = calculateFuel(fuel)

		if fuel > 0 {
			totalFuel += fuel
		}
	}

	return totalFuel
}
