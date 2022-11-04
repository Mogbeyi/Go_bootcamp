package main

import "fmt"

func area(shape string, value ...int) int {
	area := 1

	for _, val := range value {
		if shape == "square" {
			area = val * val
		} else if shape == "rectangle" {
			area *= val 
		}
	
	return area
}

func main() {
	fmt.Println(area("square", 4))
}
