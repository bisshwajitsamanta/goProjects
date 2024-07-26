package main

import "fmt"

func MergeMap(heights map[int][]string, ages map[int][]string) map[string][]int {

	// value of heights is my key for the new map
	// If the value exists in the next loop capture the value of the key as slice
	newMap := make(map[string][]int)
	for height, users := range heights {
		for _, user := range users {
			newMap[user] = append(newMap[user], height)
		}
	}

	// My first Map = map[Antoine:[180] Clement:[175] Tim:[175]]
	// Goal is if the key matches with the first map then append the second slice key to this map

	for age, users := range ages {
		for _, user := range users {
			newMap[user] = append(newMap[user], age)
		}
	}

	return newMap
}

func main() {
	fmt.Println(MergeMap(map[int][]string{175: {"Tim", "Clement"}, 180: {"Antoine"}}, map[int][]string{27: {"Clement", "Antoine"}, 21: {"Tim"}}))
}
