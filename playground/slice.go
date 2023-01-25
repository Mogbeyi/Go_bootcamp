package main

import "fmt"

func main() {
	myArray := [4]string{"First", "Second", "Third", "fourth"}
	mySlice1 := myArray
	mySlice2 := myArray
	mySlice1[0] = "Test"

	fmt.Println(myArray, mySlice1, mySlice2)
}
