package main

import (
	"fmt"
	"math"
)

func main() {
	a := 20.7
	b := math.Floor(a)
	roundOff := math.Ceil(a)
	fmt.Printf("Before Number: %.2f \t Round Off Number: %.2f", b, roundOff)
}
