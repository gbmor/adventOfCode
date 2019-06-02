package main

import (
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

func reduce(m string) string {

	buf := []rune{}
	i := 0
	for i < len(m)-2 {
		r := rune(m[i+1])
		e := rune(m[i])
		if e == unicode.ToLower(r) && unicode.IsUpper(r) {
			i++
			continue
		}
		if e == unicode.ToUpper(r) && unicode.IsLower(r) {
			i++
			continue
		}
		buf = append(buf, rune(e))
		i++
	}

	if string(buf) == m {
		return string(buf)
	}

	return reduce(string(buf))
}

func main() {
	m := getData()
	out := reduce(m.parsed)
	fmt.Printf("Reduced polymer: %v\n", out)
}
