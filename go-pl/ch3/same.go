package main

import (
	"fmt"
	"strings"
)

// 编写一个函数，判断两个字符串是否是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。

func isSame(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	bIndexMap := make(map[int]int)
	for _, i := range a {
		index := strings.Index(b, string(i))
		if index == -1 {
			return false
		}
		if _, ok := bIndexMap[index]; ok {
			return false
		}
		bIndexMap[index] = 1
	}
	return true
}

func main() {
	fmt.Println(isSame("abcde", "edacb"))
	fmt.Println(isSame("abcde", "edacaadb"))
}
