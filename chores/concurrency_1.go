package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(2)
		go func(number int) {
			defer wg.Done()
			fmt.Println("Go routine:", number)
		}(i)
	}
	wg.Wait()
}
