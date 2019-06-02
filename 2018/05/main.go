package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"unicode"
)

type material struct {
	raw     []byte
	parsed  string
	reduced string
}

func newMaterial() *material {
	return &material{
		raw:     make([]byte, 0),
		parsed:  "",
		reduced: "",
	}
}

func getData() *material {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	polymer := newMaterial()
	polymer.raw = b
	polymer.parsed = string(b)

	return polymer
}

func reduce(m *material) *material {

	var buf bytes.Buffer
	iter := m.parsed
	if m.reduced != "" {
		iter = m.reduced
	}
	flag1 := false
	flag2 := false
	for i, e := range iter {
		if i < len(m.parsed)-1 {
			if unicode.IsLower(e) && unicode.IsUpper(rune(m.parsed[i+1])) {
				if e == unicode.ToLower(rune(m.parsed[i+1])) {
					flag1 = true
					continue
				}
			}
			if unicode.IsUpper(e) && unicode.IsLower(rune(m.parsed[i+1])) {
				if e == unicode.ToUpper(rune(m.parsed[i+1])) {
					flag2 = true
					continue
				}
			}
			buf.WriteRune(e)
		}
	}
	m.reduced = buf.String()

	if !flag1 && !flag2 {
		return m
	}
	return reduce(m)
}

func main() {
	m := getData()
	reduce(m)
	fmt.Printf("Reduced polymer: %v\n", m.reduced)
}
