package main

import (
	"fmt"
	"math"
)

func matchingCharacters(string1, string2 string) int {
	count := 0
	MinimumLength := math.Min(float64(len(string1)), float64(len(string2)))

	for i := 0; i < int(MinimumLength); i++ {
		char1 := string1[i]
		char2 := string2[i]
		if char1 == char2 {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(matchingCharacters("AlgoExpert", "Alligator"))

}
