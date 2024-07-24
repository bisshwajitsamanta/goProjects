package main

import "fmt"

func main() {
	n := 10
	for a, b := 0, 1; n > 0; a, b, n = b, a+b, n-1 {
		fmt.Println(a)
	}
}
