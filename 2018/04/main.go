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
	id           int
	times        []string
	sleep        []int
	avg          int
	sleepMinutes [][]int
}

func newGuard() *guard {
	return &guard{
		id:           0,
		times:        make([]string, 0),
		sleep:        make([]int, 0),
		avg:          0,
		sleepMinutes: make([][]int, 0),
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

	g.sleepMinutes = append(g.sleepMinutes, minutes)

	for i, e := range minutes {
		if i == 1 || i%2 > 0 {
			dur := e - minutes[i-1]
			g.sleep = append(g.sleep, dur)
		}
	}

	if len(g.sleep) > 1 {
		tmp := 0
		for _, e := range g.sleep {
			tmp += e
		}
		g.avg = tmp / len(g.sleep)
	} else if len(g.sleep) == 1 {
		g.avg = g.sleep[0]
	}

	g.times = []string{}
}

func findSleepiest(gl map[int]*guard) (int, int, int) {
	id := 0
	total := 0
	min := 0

	for k, v := range gl {
		if v.avg > total {
			total = v.avg
			id = k
		}
	}

	guard := gl[id]
	sleepMap := make(map[int]int)
	for _, e := range guard.sleepMinutes {
		for k, v := range e {
			if k == 1 || k%2 > 0 {
				begin := e[k-1]
				end := v
				for i := 0; i < end; i++ {
					sleepMap[begin+i]++
				}
			}
		}
	}

	max := 0
	for k, v := range sleepMap {
		if v > max {
			min = k
			max = v
		}
	}

	return id, total, min
}

func main() {
	times := getTimes()
	parseTimes(times)
	id, total, min := findSleepiest(guardList)
	fmt.Printf("\nGuard: %v\nAverage Sleep Time: %v\nAnswer: %v\nSleepiest Minute: %v\n", id, total, id*min, min)
}
