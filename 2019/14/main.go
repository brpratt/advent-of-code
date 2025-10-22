package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brpratt/advent-of-code/file"
)

type component struct {
	chem string
	qty  int
}

type reaction struct {
	product   component
	reactants []component
}

func parseComponent(input string) component {
	split := strings.Split(input, " ")

	chem := split[1]
	qty, err := strconv.Atoi(split[0])
	if err != nil {
		panic(err)
	}

	return component{chem, qty}
}

func parseReaction(equation string) reaction {
	outer := strings.Split(equation, "=>")
	left, right := outer[0], outer[1]

	var r reaction

	r.product = parseComponent(strings.TrimSpace(right))

	for _, inner := range strings.Split(strings.TrimSpace(left), ",") {
		r.reactants = append(r.reactants, parseComponent(strings.TrimSpace(inner)))
	}

	return r
}

func scale(chem string, qty int, reactions map[string]reaction) int {
	r := reactions[chem]

	if qty <= r.product.qty {
		return 1
	}

	if qty%r.product.qty == 0 {
		return qty / r.product.qty
	}

	return (qty / r.product.qty) + 1
}

func reduced(cs map[string]int) bool {
	for k, v := range cs {
		if k != "ORE" && v > 0 {
			return false
		}
	}

	return true
}

func reduce(cs map[string]int, extra map[string]int, reactions map[string]reaction) {
	for chem, qty := range cs {
		if chem == "ORE" {
			continue
		}

		extraqty := extra[chem]
		if extraqty != 0 {
			if extraqty > qty {
				qty = 0
				cs[chem] = 0
				extraqty -= qty
				extra[chem] = extraqty
			} else {
				qty -= extraqty
				cs[chem] = qty
				extra[chem] = 0
			}
		}

		if qty == 0 {
			continue
		}

		n := scale(chem, qty, reactions)
		extraqty = n*reactions[chem].product.qty - qty

		for _, c := range reactions[chem].reactants {
			cs[c.chem] += n * c.qty
		}

		cs[chem] = 0
		extra[chem] += extraqty
	}
}

func ore(cs map[string]int, reactions map[string]reaction) int {
	extra := make(map[string]int)

	for !reduced(cs) {
		reduce(cs, extra, reactions)
	}

	return cs["ORE"]
}

func part01(equations []string) int {
	reactions := make(map[string]reaction)

	for i := range equations {
		r := parseReaction(equations[i])
		reactions[r.product.chem] = r
	}

	cs := make(map[string]int)

	cs["FUEL"] = 1

	return ore(cs, reactions)
}

func part02(equations []string) int {
	reactions := make(map[string]reaction)

	for i := range equations {
		r := parseReaction(equations[i])
		reactions[r.product.chem] = r
	}

	fuel := 0

	for step := 10000000; step > 0; step /= 10 {
		o := 0
		for o < 1000000000000 {
			fuel += step
			cs := make(map[string]int)
			cs["FUEL"] = fuel

			o = ore(cs, reactions)
		}

		fuel -= step
	}

	return fuel
}

func main() {
	equations := file.Must(file.ReadLines("input.txt"))

	fmt.Println(part01(equations))
	fmt.Println(part02(equations))
}
