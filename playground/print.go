package main

import "fmt"

func greet(name string) {
	fmt.Printf("zola %s \n", name)
}

func main() {
	friends := []string{"Dammy", "Femi", "Ife", "Opyma"}

	for _, friend := range friends {
		greet(friend)
	}

}
