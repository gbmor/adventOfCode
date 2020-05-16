package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func getInput() []int {
	raw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	rawStr := string(raw)
	rawLines := strings.Split(rawStr, "\n")
	ints := []int{}

	for _, e := range rawLines {
		if e == "" {
			continue
		}

		parsed, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}

		ints = append(ints, parsed)
	}

	return ints
}

func sum(ints []int) int {
	out := 0
	for _, e := range ints {
		out += e
	}
	return out
}

func naiveFuel(ints []int) int {
	out := 0

	for _, e := range ints {
		out += int(math.Floor(float64(e)/3)) - 2
	}

	return out
}

func allFuel(ints []int) int {
	list := []int{}

	for _, e := range ints {
		tmp := 0
		for e > 0 {
			e = int(math.Floor(float64(e)/3)) - 2
			if e > 0 {
				tmp += e
			}
		}
		list = append(list, tmp)
	}

	return sum(list)
}

func main() {
	input := getInput()
	naive := naiveFuel(input)
	fmt.Printf("Part 1 Fuel: %d\n", naive)
	all := allFuel(input)
	fmt.Printf("Part 2 Fuel: %d\n", all)
}
