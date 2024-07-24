package main

import "fmt"

var stop bool

func main() {
	i := 0
	for !stop { // While tell stop not to True
		fmt.Println("Hello Bisshwajit!!")
		i++
		if i == 10 {
			stop = true
		}
	}

}
