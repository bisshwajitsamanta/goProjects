package main

import "fmt"

func main() {
	for i, j := 0, 100; i < 10; i, j = i+1, j-10 {
		fmt.Printf("i:%d,j:%d\n", i, j)
	}
}
