package main

import (
	"fmt"
	"math"
)

func RoundedSquareRoot(number float64) float64 {
	var output float64
	if number > 0 {
		output = math.Round(math.Sqrt(number))

	}
	return output
}

func main() {
	fmt.Println(RoundedSquareRoot(11))
}
