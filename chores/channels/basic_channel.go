package main

import "fmt"

func main() {
	numbersCh := make(chan int)

	go func() {
		numbersCh <- 10
	}()
	fmt.Println(<-numbersCh)
}
