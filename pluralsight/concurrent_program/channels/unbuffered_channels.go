package main

import "sync"

var (
	ch = make(chan string)
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- "message"
	}()
	go func() {
		defer wg.Done()
		//msg := <-ch
		//print(msg)
		print(<-ch)
	}()
	wg.Wait()
}
