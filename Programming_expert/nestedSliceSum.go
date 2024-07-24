package main

import "fmt"

func NestedSliceSum(numbers [][]int) []int {
	// Write your code here.
	var sums []int
	for i, nestedValue := range numbers {
		sums = append(sums, 0)
		fmt.Println("Sums Array:", sums)
		for _, value := range nestedValue {
			sums[i] += value

		}
	}
	return sums
}

func main() {
	fmt.Println(NestedSliceSum([][]int{{3, 4, 5}, {-3, -4, -5}, {-1, -1, 0}, {-4, 25}}))
}
