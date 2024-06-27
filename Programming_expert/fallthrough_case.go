package main

import "fmt"

func main() {
	switch a := 0; {
	//case a < 1:
	case true:
		fmt.Println("One")
		fallthrough
	case a < 2:
		fmt.Println("two")

	}
}
