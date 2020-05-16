package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput() []int {
	raw, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rawStr := string(raw)
	rawFields := strings.Split(rawStr, ",")
	out := []int{}

	for _, e := range rawFields {
		e = strings.TrimSpace(e)
		if e == "" {
			continue
		}
		i, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}

	return out
}

func computePos0(ints []int, noun, verb int) int {
	ints[1] = noun
	ints[2] = verb

	for i := 0; i < len(ints)-3; i += 4 {
		pos1 := ints[i+1]
		pos2 := ints[i+2]
		dest := ints[i+3]

		switch ints[i] {
		case 1:
			ints[dest] = ints[pos1] + ints[pos2]
		case 2:
			ints[dest] = ints[pos1] * ints[pos2]
		case 99:
			return ints[0]
		}
	}

	return ints[0]
}

func bruteforce() int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if computePos0(getInput(), noun, verb) == 19690720 {
				return 100 * noun + verb
			}
		}
	}

	return 0
}

func main() {
	input := getInput()
	pos0 := computePos0(input, 12, 2)
	fmt.Printf("(12, 2) Position 0: %d\n", pos0)
	fmt.Printf("100 * Noun + Verb for Pos0 == 19690720: %d\n", bruteforce())
}
