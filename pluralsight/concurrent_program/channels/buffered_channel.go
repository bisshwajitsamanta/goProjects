package main

import (
	"fmt"
	"sync"
)

var channel = make(chan string, 1)

func sendMessage(wg *sync.WaitGroup) {
	defer wg.Done()
	channel <- "Hi Bisshwajit !!"
}

func ReceiveMessage(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-channel)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go sendMessage(&wg)
	go ReceiveMessage(&wg)
	wg.Wait()

}
