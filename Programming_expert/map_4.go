package main

import "fmt"

func main() {
	n := 11
	mp := map[int]uint{}
	for i := 1; i <= n; i++ {
		for d := 1; d <= 5; d++ {
			if i%d == 0 {
				mp[d]++
			}
		}
	}
	fmt.Println(mp)
}
