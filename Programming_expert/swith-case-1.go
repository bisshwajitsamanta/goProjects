package main

import "fmt"

func main() {
	a := 3

	/* switch a here is comparing the values
	switch a ==1 then going to case 1
	switch a ==2 then going to case 2
	Here after case keyword it is expecting a value like 1,2,3 etc.
	*/
	switch a {
	case 1:
		fmt.Println("The value is one")
	case 2:
		fmt.Println("The value is two")
	default:
		fmt.Println("You have reached to default !!")
	}
}
