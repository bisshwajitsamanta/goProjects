package main

import (
	"fmt"
	"sync"
)

var newCh = make(chan string, 1)

func WrapperWaitGroup(wg *sync.WaitGroup, fn func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fn()
	}()
}

func send_Message() {
	newCh <- "Hello Bisshwajit !!"
}

func receive_Message() {
	fmt.Println(<-newCh)
}

func main() {
	wg := sync.WaitGroup{}
	WrapperWaitGroup(&wg, send_Message)
	WrapperWaitGroup(&wg, receive_Message)
	wg.Wait()
}
