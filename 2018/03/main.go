package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type fabric struct {
	grid [][]int
}

func newFabric() *fabric {
	return &fabric{
		grid: [][]int{},
	}
}

func (f *fabric) draw() {

}

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

// Columns of input are as follows:
// 0: Claim number
// 1: Inches from left side of fabric
// 2: Inches from top of fabric
// 3: Width (inches) of claim
// 4: Height (inches) of claim
func parseClaims() [][]int {
	var out = [][]int{}
	input, err := ioutil.ReadFile("input.txt")
	maybePanic(err)
	lines := strings.Split(string(input), "\n")
	for k, v := range lines {
		if v == "" {
			continue
		}
		out = append(out, []int{})
		cols := strings.Split(v, "\t")

		claim, err := strconv.Atoi(cols[0])
		maybePanic(err)
		out[k] = append(out[k], claim)

		x, err := strconv.Atoi(cols[1])
		maybePanic(err)
		y, err := strconv.Atoi(cols[2])
		maybePanic(err)
		out[k] = append(out[k], x)
		out[k] = append(out[k], y)

		fillx, err := strconv.Atoi(cols[3])
		maybePanic(err)
		filly, err := strconv.Atoi(cols[4])
		maybePanic(err)

		out[k] = append(out[k], fillx)
		out[k] = append(out[k], filly)
	}

	return out
}

func main() {
	out := parseClaims()
	fmt.Printf("%#v\n", out)
}
