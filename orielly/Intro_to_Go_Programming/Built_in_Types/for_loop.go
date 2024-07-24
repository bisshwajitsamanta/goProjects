package main

import "fmt"

func main() {
	for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
		fmt.Printf("%d %d Hello world\n", i, j)
	}
}
