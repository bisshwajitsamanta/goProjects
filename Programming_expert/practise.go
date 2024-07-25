/*
	Practise Questions:
		1. Fibonacci Series
		2. Nested Slice
*/

package main

import "fmt"

func NestedSlice(numbers [][]int) []int {
	sums := make([]int, len(numbers))

	for i, nestedValue := range numbers {
		for _, value := range nestedValue {
			sums[i] += value
		}
	}

	return sums
}

func main() {
	fmt.Println(NestedSlice([][]int{{3, 4, 5}, {5, -7, -8}, {4, 25}}))
}
