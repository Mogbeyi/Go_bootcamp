// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	newHeader := strings.Split(header, ",")
	newData := strings.Split(data, "\n")

	for i := range newHeader {
		fmt.Printf("%-15s", newHeader[i])
	}
	fmt.Println()
	fmt.Println("===========================================================================")

	for i := range newData {
		splitData := strings.Split(newData[i], ",")
		for _, value := range splitData {
			fmt.Printf("%-15s", value)
		}
		fmt.Println()
	}

	fmt.Println(newData)

	var (
		size, bed, bath, price int
	)

	for i := range newData {
		splitData := strings.Split(newData[i], ",")
		intSize, _ := convertToInt(splitData[1])
		intBed, _ := convertToInt(splitData[2])
		size += intSize
		bed += intBed
		intBath, _ := convertToInt(splitData[3])
		bath += intBath
		intPrice, _ := convertToInt(splitData[4])
		price += intPrice
	}

	fmt.Printf("%-12s %-15f %-15f %-15f %-15f", "", float64(size / 4), float64(bed / 4), float64(bath / 4), float64(price / 4))
}

func convertToInt(n string) (int, error) {
	value, err := strconv.Atoi(n)
	return value, err 
}