package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return 
	}

	user, password := args[1], args[2]

	if user != "emmanuel" || password != "1234" {
		fmt.Println("Wrong username or password")
	} else {
		fmt.Printf("Welcome %s", user)
	}
}