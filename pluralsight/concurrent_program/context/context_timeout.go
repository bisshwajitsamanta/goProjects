package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for range time.Tick(time.Second) {
			if ctx.Err() != nil {
				log.Println(ctx.Err())
				return
			}
			fmt.Println("Tick Tick!!")
		}
	}(ctx)
	wg.Wait()
}
