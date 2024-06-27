package main

import "fmt"

func Factor(n int) int {

	// 5!=5*4*3*2*1
	value := 1
	for {
		if n > 0 {
			value *= n
			n--

		} else {
			break

		}
	}
	return value
}

func main() {
	fmt.Println(Factor(5))
}
