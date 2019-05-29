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
	for i, e := range boxen {
		for k, v := range boxen {
			if i == k {
				continue
			}
			samesies := []rune{}
			diffs := 0
			compstr := []rune(v)
			for j, p := range e {
				if p != compstr[j] {
					diffs++
				}
				if diffs > 1 {
					break
				}
				if p == compstr[j] {
					samesies = append(samesies, p)
				}
				if j == len(e)-1 {
					return string(samesies)
				}
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
