package main

import "fmt"

func setItem(slice []int, index, value int) []int {
	slice[index] = value

	return slice
}

func main() {
	slice := []int{1, 2, 3, 4}
	newSlice := setItem(slice, 3, 6)

	fmt.Println(newSlice)
}
