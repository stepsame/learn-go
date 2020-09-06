package main

import "math"

// get max number
func Max(n ...int) int {
	if len(n) == 0 {
		return 0
	}
	maxV := math.MinInt64
	for _, v := range n {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}

func main() {
	println(Max(1, 2, 3, 4, 9))
	println(Max())
}
