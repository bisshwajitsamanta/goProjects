package main

import (
	"strconv"
)

func MultiplyByString(str string, number int) int {
	// Write your code here.

	output, _ := strconv.Atoi(str)

	return output * number
}
