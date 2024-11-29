package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	for _, v := range []string{"Bisshwajit", "Samanta", "sinchan"} {
		ch <- v
	}
	close(ch)
	for msg := range ch {
		fmt.Println(msg)
	}
}
