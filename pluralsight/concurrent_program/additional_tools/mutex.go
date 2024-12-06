package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{}
	var wg sync.WaitGroup
	var m sync.Mutex
	const iterations = 10000
	wg.Add(iterations)
	for i := range iterations {
		go func() {
			defer wg.Done()
			m.Lock()
			defer m.Unlock()
			s = append(s, i)
		}()
	}
	wg.Wait()
	fmt.Println(len(s))
}
