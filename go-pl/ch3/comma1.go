package main

import (
	"fmt"
	"bytes"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
    n := len(s)
	var buf bytes.Buffer
    if n <= 3 {
        return s
    }

    prefix := strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-")
    first := 0
    if prefix {
    	first = 1
	}
    point := strings.LastIndex(s, ".")
    if point == -1 {
    	point = n - 1
	}
	for i, c := range s[first:point] {
		buf.WriteRune(c)
		if i % 3 == 2 {
			buf.WriteByte(',')
		}
	}
    return s[:first] + buf.String() + s[point:]
}

func main() {
	fmt.Println(comma("12345123.123132"))
	fmt.Println(comma("12345123"))
	fmt.Println(comma("+12345123"))
}
