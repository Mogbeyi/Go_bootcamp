package main

import "fmt"

func main() {
	type Empty interface {}

	var e Empty
	fmt.Printf("e's value: %v, type: %T\n", e, e)

	e = 7
	fmt.Printf("e's value: %v, type: %T\n", e, e)

	e = "Hello World"
	fmt.Printf("e's value: %v, type: %T\n", e, e)
}