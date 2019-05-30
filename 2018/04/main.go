package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var guardList = make(map[int]*guard)

type guard struct {
	id    int
	times []string
	sleep int
}

func newGuard() *guard {
	return &guard{
		id:    0,
		times: make([]string, 0),
		sleep: 0,
	}
}

func quickErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getTimes() []string {
	var times []string

	out, err := ioutil.ReadFile("input-sorted.txt")
	quickErr(err)
	scanner := bufio.NewScanner(bytes.NewReader(out))

	for scanner.Scan() {
		times = append(times, scanner.Text())
	}

	return times
}

func parseTimes(times []string) {
	lastGuard := 0

	for i, e := range times {
		cols := strings.Split(e, " ")
		timestamp := cols[1]

		if strings.HasPrefix(cols[3], "#") {
			if i > 0 {
				guardList[lastGuard].sleepTotal()
			}
			col := cols[3]
			id, err := strconv.Atoi(col[1:])
			quickErr(err)
			lastGuard = id
			if _, ok := guardList[id]; !ok {
				guard := newGuard()
				guard.id = id
				guardList[id] = guard
			}
			continue
		}
		guard := guardList[lastGuard]
		guard.times = append(guard.times, timestamp)
	}
}

func (g *guard) sleepTotal() {
	var minutes []int
	for _, e := range g.times {
		split := strings.Split(e, ":")
		mins := split[1]

		num, err := strconv.Atoi(mins)
		quickErr(err)

		minutes = append(minutes, num)
	}

	for i, e := range minutes {
		if i == 1 || i%2 > 0 {
			dur := e - minutes[i-1]
			g.sleep += dur - 1
		}
	}

	fmt.Printf("%#v\n", g)
	g.times = []string{}
	fmt.Printf("%#v\n", g)

}

func findSleepiest(gl map[int]*guard) (int, int) {
	id := 0
	total := 0

	for k, v := range gl {
		if v.sleep > total {
			total = v.sleep
			id = k
		}
	}

	return id, total
}

func main() {
	times := getTimes()
	parseTimes(times)
	id, mins := findSleepiest(guardList)
	fmt.Printf("\nGuard: %v\nMinutes Asleep: %v\nAnswer: %v\n", id, mins, id*mins)
}
