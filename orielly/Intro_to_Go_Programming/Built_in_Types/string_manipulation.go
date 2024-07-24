package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "hello"
	output := make([]string, len(input))
	for i, ch := 0, 'a'; i < len(input); i, ch = i+1, ch+1 {
		output[i] = strings.ToUpper(string(ch))
	}
	fmt.Println(strings.Join(output, ""))
}
