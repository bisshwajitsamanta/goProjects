package main

import "fmt"

func main() {
	str := "hello"
	for _, v := range str {
		fmt.Println(string(v))
	}
}
