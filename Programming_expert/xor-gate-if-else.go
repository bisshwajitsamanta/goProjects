package main

import "fmt"

func XOR(x bool, y bool) bool {
	// Write your code here.
	if x == true && y == false {
		return true
	} else if x == false && y == false {
		return false
	} else if x == false && y == true {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(XOR(true, false))
}
