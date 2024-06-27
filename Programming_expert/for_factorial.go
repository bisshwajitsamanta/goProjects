package main

import "fmt"

func Factorial(n int) int {
	// Write your code here.
	/*
		n >0 first, so the initial value is 1
		Decrement the input to n to 1 and store somewhere and multiply among each other

	*/
	result := 1
	for j := n; j >= 1; j-- {
		result *= j
	}

	return result
}

func main() {
	fmt.Println(Factorial(5))
}
