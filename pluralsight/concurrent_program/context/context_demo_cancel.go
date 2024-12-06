package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Context Package Demo")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for range time.Tick(time.Second) {
			print(ctx.Err())
			if ctx.Err() != nil {
				fmt.Println("Context was cancelled", ctx)
				return
			}
			fmt.Println("Tick !!")
		}
	}(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}
