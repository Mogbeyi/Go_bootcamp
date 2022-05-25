package main

import (
	"fmt"
	"strconv"
	"os"
)


/*
if year is divisible by 400 then is_leap_year
else if year is divisible by 100 then not_leap_year
else if year is divisible by 4 then is_leap_year
else not_leap_year
*/

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Entered wrong number of command line arguments")
	}
	arg := os.Args[1]
	year, err := strconv.Atoi(arg)

	if err != nil {
		fmt.Printf("%v is not an integer", arg)
		return 
	}

	if year % 400 == 0 {
		fmt.Printf("%d is a leap year", year)
	} else if year % 100 == 0 {
		fmt.Printf("%d is not a leap year", year)
	} else if year % 4 == 0 {
		fmt.Printf("%d is a leap year", year)
	} else {
		fmt.Printf("%d is not a leap year", year)
	}
}