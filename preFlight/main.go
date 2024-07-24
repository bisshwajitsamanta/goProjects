package main

import (
	"fmt"
	"os/exec"
)

func commandCheck() string {

	//services := []string{"docker", "nginx", "smb", "zfs", "auditd"}
	services := []string{"docker", "wc", "date"}
	var result string
	for _, v := range services {
		_, err := exec.LookPath(v)
		if err != nil {
			result += fmt.Sprintf("Cannot find %s in our PATH\n", v)
		}
		result += fmt.Sprintf("Found %s in our PATH\n", v)
	}

	return result
}

func main() {
	fmt.Println("Hello")
	go commandCheck()

}
