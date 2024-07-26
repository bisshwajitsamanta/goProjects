package main

import "fmt"

func main() {
	mp := map[string]int{"h": 1, "a": 2, "c": 8}
	for k, v := range mp {
		fmt.Println(k, v)
	}
}
