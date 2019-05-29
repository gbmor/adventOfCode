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
		prev := []rune(boxen[i-1])
		cur := []rune(boxen[i])
		samesies = []rune{}

		for k, v := range prev {
			if v != cur[k] {
				diffs++
				if diffs > 1 {
					break
				}
			}
			if v == cur[k] {
				samesies = append(samesies, v)
			}
			if k == len(prev)-1 && diffs <= 1 {
				return string(samesies)
			}
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
