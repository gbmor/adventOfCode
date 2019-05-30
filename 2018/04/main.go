package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type guard struct {
	id       int
	schedule map[time.Time]string
}

func newGuard() *guard {
	return &guard{
		id:       0,
		schedule: make(map[time.Time]string),
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

func parseTimes(times []string) map[int]*guard {
	guards := make(map[int]*guard)
	lastGuard := 0

	for _, e := range times {
		cols := strings.Split(e, " ")
		timestamp := fmt.Sprintf("%v %v", cols[0], cols[1])
		thetime, err := time.Parse("2019-04-29 13:00", timestamp)
		quickErr(err)

		var guard *guard
		if strings.HasPrefix(cols[3], "#") {
			var err error
			guard = newGuard()
			col := cols[3]
			guard.id, err = strconv.Atoi(col[1:])
			quickErr(err)
			lastGuard = guard.id
			guard.schedule[thetime] = "start"
			continue
		}
		guard = guards[lastGuard]
		if strings.Contains(cols[2], "falls") {
			guard.schedule[thetime] = "sleep"
			continue
		}
		if strings.Contains(cols[2], "wakes") {
			guard.schedule[thetime] = "wake"
		}

	}

	return guards
}
