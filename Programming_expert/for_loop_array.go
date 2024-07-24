package main

import "fmt"

func main() {
	arr2 := [2][2]string{{"hello", "Bisshwajit"}, {"go", "code"}}
	fmt.Println(arr2)
	fmt.Printf("Type of array: %T\n", arr2)
	// Looping over the array
	for _, nestedArray := range arr2 {
		for j, str := range nestedArray {
			fmt.Println(j, str)
		}
	}
}
