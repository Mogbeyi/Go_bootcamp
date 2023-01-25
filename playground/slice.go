package main

import "fmt"

func main() {
	myArray := [3]string{"First", "Second", "Third"}
	mySlice1 := myArray
	mySlice2 := myArray
	mySlice1[0] = "Test"

	fmt.Println(myArray, mySlice1, mySlice2)
}
