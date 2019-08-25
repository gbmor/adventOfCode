package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func finder(box string) (bool, bool) {
	two, three := false, false
	letmap := make(map[rune]int)

	for _, e := range box {
		letmap[e]++
	}
	for _, v := range letmap {
		if v == 2 {
			two = true
		}
		if v == 3 {
			three = true
		}
	}

	return two, three
}

func checksum(boxen []string) int {
	var one, two int

	for _, e := range boxen {
		a, b := finder(e)
		if a {
			one++
		}
		if b {
			two++
		}
	}

	return one * two
}

func hammingDistance(a, b string) int {
	var dist int
	for k, v := range a {
		if v != rune(b[k]) {
			dist++
		}
	}
	return dist
}

func oneCharDiff(boxen []string) string {
	var hams = map[int][]string{}
	for _, v := range boxen {
		for _, e := range boxen {
			dist := hammingDistance(e, v)
			hams[dist] = []string{e, v}
		}
	}

	var two []string
	for k, v := range hams {
		if k == 1 {
			two = v
		}
	}

	var samesies []rune
	for i, e := range two[0] {
		r := []rune(two[1])
		if e == r[i] {
			samesies = append(samesies, e)
		}
	}

	return string(samesies)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Printf("%v\n", err)
	}

	buf := bytes.NewBuffer(data)
	bs := bufio.NewScanner(buf)
	var boxen []string

	for bs.Scan() {
		boxen = append(boxen, bs.Text())
	}

	cksum := checksum(boxen)
	diff := oneCharDiff(boxen)

	fmt.Printf("Checksum: %v\n", cksum)
	fmt.Printf("Matching box code: %v\n", diff)
}
