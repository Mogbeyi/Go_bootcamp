package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	
	if len(args) != 3 {
		fmt.Println("Usage:[username] [password]")
	}

	username, password := args[1], args[2]

	if username == "emmanuel" && password == "1234" {
		fmt.Printf("Welcome %s", username)
	} else if username == "damilare" && password == "6789" {
		fmt.Printf("Welcome %s", username)
	} else {
		fmt.Println("Wrong username or password")
	}
}