package main

import "fmt"

func main() {
	a := "h"
	switch a {
	case "h", "a", "b":
		fmt.Println("worked")
	default:
		fmt.Println("You have reached Dead End !!")

	}
}
