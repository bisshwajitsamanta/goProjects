package main

import "fmt"

func main() {
	mp := make(map[int]rune)
	mp[5] = 2
	mp[6] = 4
	delete(mp, 6)
	value, ok := mp[5]
	fmt.Println(value, ok)
	fmt.Println(mp)
}
