package main

import "fmt"

func main() {
	// Appending to a slice
	var arr []string
	//arr := []string{}
	for _ = range 10 {
		arr = append(arr, "Hello")
		fmt.Println(arr, len(arr), cap(arr))
	}
}
