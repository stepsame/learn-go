package main

import "fmt"

func rotate(s []int, i int) []int {
	res := make([]int, len(s))
	for index, v := range(s) {
		if index >= i {
			res[index -i] = v
		} else {
			res[len(s)-i+index] = v
		}
	}
	return res

}

func main()  {
	// 0 1 2 3 4
	// 3 4 0 1 2
	a := []int{1,2,3,4,5}
	fmt.Println(rotate(a, 2))
}
