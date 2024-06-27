package main

import "fmt"

func VerifyLogin(username string, password string, attempt int8) bool {
	// Write your code here.
	if username == "ProgrammingExpert" && password == "LearnToCode" && attempt < 5 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(VerifyLogin("ProgrammingExpert", "LearnToCode", 1))
}
