package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sliceArr := arr[:4]
	fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))
	sliceArr = sliceArr[1:]
	fmt.Println(sliceArr, len(sliceArr), cap(sliceArr))
}
