package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type fabric struct {
	grid [1000][1000]int
}

func newFabric() *fabric {
	return &fabric{
		grid: [1000][1000]int{},
	}
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

func (f *fabric) drawClaims(claims [][]int) {
	for _, v := range claims {
		x := v[1]
		y := v[2]
		for i := 0; i < v[3]; i++ {
			f.grid[y][x+i]++
			for j := 1; j < v[4]; j++ {
				f.grid[y+j][x+i]++
			}
		}
	}
}

func (f *fabric) findOverlap() int {
	var out int

	for _, v := range f.grid {
		for _, e := range v {
			if e > 1 {
				out++
			}
		}
	}

	return out
}

func main() {
	fab := newFabric()
	claims := parseClaims()
	fab.drawClaims(claims)
	fmt.Printf("Overlap: %v square inches\n", fab.findOverlap())

}
