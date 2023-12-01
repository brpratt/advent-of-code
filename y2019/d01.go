package y2019

import (
	"bufio"
	"io"
	"strconv"
)

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

func readNumbers(r io.Reader) []int {
	masses := make([]int, 0)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		masses = append(masses, mass)
	}

	return masses
}

func SolveD01P01(r io.Reader) (string, error) {
	input := readNumbers(r)
	fuel := 0

	for _, mass := range input {
		fuel += calculateFuel(mass)
	}

	return strconv.Itoa(fuel), nil
}

func SolveD01P02(r io.Reader) (string, error) {
	input := readNumbers(r)
	fuel := 0

	for _, mass := range input {
		fuel += calculateFuelWithOverhead(mass)
	}

	return strconv.Itoa(fuel), nil
}
