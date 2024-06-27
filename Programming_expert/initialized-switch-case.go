package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("one")
	case true:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}
}
