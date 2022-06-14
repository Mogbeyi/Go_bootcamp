package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	num, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Printf("%v is not a number", arg)
		return 
	}

	if num % 2 != 0 {
		fmt.Printf("%d is an odd number", num)
	} else if num % 2 == 0 && num % 8 == 0 {
		fmt.Printf("%d is an even number and is divisible by 8", num)
	} else {
		fmt.Printf("%d is an even number", num)
	}
}