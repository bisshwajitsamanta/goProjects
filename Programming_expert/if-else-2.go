package main

import "fmt"

func ColorMaker(color1 string, color2 string) string {
	// Write your code here.
	if color1 == "red" && color2 == "yellow" {
		return "orange"
	} else if color1 == "yellow" && color2 == "red" {
		return "orange"
	} else if color1 == "yellow" && color2 == "blue" {
		return "green"
	} else if color1 == "blue" && color2 == "yellow" {
		return "green"
	} else if color1 == "red" && color2 == "red" {
		return "red"
	} else if color1 == "yellow" && color2 == "yellow" {
		return "yellow"
	} else if color1 == "blue" && color2 == "blue" {
		return "blue"
	} else if color1 == "blue" && color2 == "red" {
		return "violet"
	} else {
		return "violet"
	}
}

func main() {
	fmt.Println(ColorMaker("blue", "yellow"))
}
