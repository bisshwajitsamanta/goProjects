package main

import "fmt"

func CharacterFrequency(sentence string) map[rune]int {
	// Write your code here.
	characters := map[rune]int{}
	for _, char := range sentence {
		// characters[char]++
		if value, ok := characters[char]; ok {
			characters[char] = value + 1
		} else {
			characters[char] = 1
		}
	}

	return characters
}

func main() {
	fmt.Println(CharacterFrequency("hello world"))
}
