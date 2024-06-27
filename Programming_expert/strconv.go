package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "1234"
	b, err := strconv.Atoi(a)
	fmt.Printf("Value of b: %d \nValue of err: %v\n", b, err)

	str := "2345"
	c, err := strconv.ParseInt(str, 10, 64)
	fmt.Printf("Value of c: %d\nError: %v", c, err)
}
