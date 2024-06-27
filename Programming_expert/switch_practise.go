package main

import "fmt"

func DailyRoutine(time string) string {
	// Write your code here.
	switch time {
	case "morning":
		return "Time to Get up"
	case "afternoon":
		return "Time for Lunch"
	case "night":
		return "Time to sleep!"
	default:
		return "Have a great day :) "
	}
}

func main() {
	fmt.Println(DailyRoutine("morning"))
}
