package main

import "fmt"

func factorial(n int) int {
	value := 1
	for i := 1; i <= n; i++ {
		value *= i
	}
	return value
}

func main() {
	fmt.Println(factorial(5))
}
