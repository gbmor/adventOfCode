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

func oneCharDiff(boxen []string) string {
	var samesies []rune
	for i := 1; i < len(boxen); i++ {
		diffs := 0
		for p, e := range boxen[i] {
			if diffs > 1 {
				samesies = []rune{}
				break
			}
			prev := boxen[i-1]
			if e == rune(prev[p]) {
				samesies = append(samesies, e)
			} else if e == rune(prev[p]) && p == len(boxen[i])-1 {
				samesies = append(samesies, e)
				return string(samesies)
			} else {
				diffs++
			}
		}
		if diffs == 1 {
			break
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

	fmt.Printf("Checksum: %v\n", cksum)

	fmt.Printf("Matching box code: %v\n", oneCharDiff(boxen))
}
