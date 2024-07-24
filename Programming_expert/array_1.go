package main

import "fmt"

func main() {
	var arr [2]int
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Printf("Length of Array: %d", len(arr))
	newArr := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(newArr)
	fmt.Println(len(newArr))
}
