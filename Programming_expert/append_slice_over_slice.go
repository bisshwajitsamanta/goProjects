package main

import "fmt"

func main() {
	arr := []bool{}
	arr2 := []bool{true, false, true}
	arr = append(arr, arr2...) // Append 1 slice over another ... ( ellipsis )
	fmt.Println(arr)
}
